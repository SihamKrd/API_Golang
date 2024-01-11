import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

#from src.schemas.rating import RatingSchema
from src.models.http_exceptions import *


ratings_url = "http://ratings-lima.edu.forestier.re/songs/"   # URL de l'API ratings


def get_rating(id,songID):
    url = ratings_url+ songID+"/ratings/"+id
    response = requests.request(method="GET", url=url)
    return response.json(), response.status_code


def get_all_ratings(songID):
    url = ratings_url+ songID+"/ratings"
    response = requests.request(method="GET", url=url)
    return response.json(), response.status_code


def create_rating(rating,songID):
    # on crée le commentaire
    rating['user_id'] = current_user.id
    rating['song_id'] = songID
    response = requests.request(method="POST", url=ratings_url+ songID+"/ratings", json=rating)
    
    if response.status_code != 201:
        return response.json(), response.status_code
    return response.json(), response.status_code



def update_rating(songID,id, rating_update):

    rating_update['user_id'] = current_user.id
    rating_update['song_id'] = songID
    rating_update['id']=id
    print(rating_update['user_id'])
    print("est ce que je rentre ici")
        # on lance la requête de modification
    response = requests.request(method="PUT", url=ratings_url+ songID+"/ratings/"+id, json=rating_update)
    print (response.status_code)
    if response.status_code != 200:
        return response.json(), response.status_code
    return (response.json(), response.status_code) if response else get_rating(id,songID)



def delete_rating(songID,id):
    rating_response = requests.get(url=ratings_url + songID + "/ratings/" + id)
    response = requests.request(method="DELETE", url=ratings_url+ songID+"/ratings/"+id)
    return "", response.status_code
