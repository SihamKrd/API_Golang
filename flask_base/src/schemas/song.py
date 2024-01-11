from marshmallow import Schema, fields, validates_schema, ValidationError
               

# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")
    album = fields.String(description="Album")
    genre = fields.String(description="genre")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("artist") or obj.get("artist") == "") and \
               (not obj.get("album") or obj.get("album") == "") and \
               (not obj.get("genre") or obj.get("genre") == "")

class BaseSongSchema(Schema):
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")
    album = fields.String(description="Album")
    genre = fields.String(description="genre")


# Schéma utilisateur de modification (name, username, password)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "") or
                ("album" in data and data["album"] != "") or
                ("genre" in data and data["genre"] != "")) :
            raise ValidationError("at least one of ['title','artist','album','genre'] must be specified")
