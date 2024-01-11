import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.rating import RatingSchema
from src.schemas.rating import RatingUpdateSchema
from src.schemas.errors import *
import src.services.ratings as ratings_service

from src.helpers.content_negotiation import content_negotiation

# from routes import ratings
ratings = Blueprint(name="ratings", import_name=__name__)



# Get all ratings related to a song 
@ratings.route('/<songID>/ratings', methods=['GET'])
@content_negotiation
def get_all_ratings(songID):
    """
    ---
    get:
      description: Get a list of all ratings
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: 
                type : array
                items : ratingSchema
            application/yaml:
              schema: 
                type : array
                items : ratingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
          - ratings
    """
    return ratings_service.get_all_ratings(songID)



@ratings.route('/<songID>/ratings/<id>', methods=['GET'])
@content_negotiation
def get_rating(id,songID):
    """
    ---
    get:
      description: Getting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: RatingSchema
            application/yaml:
              schema: RatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
          - ratings
    """
    return ratings_service.get_rating(id,songID)



@ratings.route('/<songID>/ratings', methods=['POST'])
@content_negotiation
@login_required
def post_rating(songID):
    """
    ---
    post:
      description: Add a rating
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingSchema
      responses:
        '201':
          description: Rating created successfully
          content:
            application/json:
              schema: RatingSchema
            application/yaml:
              schema: RatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound 
      tags:
          - songs
          - ratings
    """
    # parser le body
    try:
        rating = RatingSchema().loads(json_data=request.data.decode('utf-8'))
        print(rating)
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # creation du rating
    try:
        return ratings_service.create_rating(rating,songID)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")




@ratings.route('/<songID>/ratings/<id>', methods=['PUT'])
@content_negotiation
@login_required
def put_rating(songID,id):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: songID
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdateSchema
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: RatingSchema
            application/yaml:
              schema: RatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
          - ratings
    """
    # parser le body
    try:
        rating_update = RatingUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification du commentaire 
    try:
        return ratings_service.update_rating(songID,id, rating_update)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other users"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")



@ratings.route('/<songID>/ratings/<id>', methods=['DELETE'])
@content_negotiation
@login_required
def delete_rating(songID,id):
    """
    ---
    delete:
      description: Delete a rating by ID
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of the rating to be deleted
      responses:
        '200':
          description: Rating deleted successfully
          content:
            application/json:
              schema: SuccessResponse  
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
          - ratings
    """
    return ratings_service.delete_rating(songID,id) 



 
