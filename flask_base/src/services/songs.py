import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE

from src.schemas.song import SongSchema
from src.models.http_exceptions import *


songs_url = "http://localhost:8088/songs/"  # URL de l'API songs (golang)


def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code


def get_all_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code


def create_song(song):
    # on récupère le schéma song pour la requête vers l'API songs
    song_schema = SongSchema().loads(json.dumps(song), unknown=EXCLUDE)

    # on crée la chanson
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code
    return response.json(), response.status_code


def modify_song(id, song_update):

    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=songs_url+id, json=song_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    return (response.json(), response.status_code) if response else get_song(id)



def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+id)
    return "", response.status_code
