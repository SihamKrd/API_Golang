from functools import wraps
from flask import request, Response
import json
import yaml

def content_negotiation(f):
    @wraps(f)
    def wrapper(*args, **kwargs):
        result = f(*args, **kwargs)

        # Gérer les cas où la fonction de vue renvoie un tuple (data, status_code)
        if isinstance(result, tuple):
            response_data, status_code = result
        else:
            # Si seulement les données sont renvoyées
            response_data = result
            status_code = 200

        # Vérifier le header 'Accept'
        if "application/x-yaml" in request.headers.get("Accept", ""):
            formatted_response = yaml.dump(response_data)
            content_type = "application/x-yaml"
        else:
            # Par défaut, utiliser JSON
            formatted_response = json.dumps(response_data)
            content_type = "application/json"

        return Response(formatted_response, status=status_code, content_type=content_type)
    return wrapper
