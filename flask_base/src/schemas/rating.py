from marshmallow import Schema, fields, validates_schema, ValidationError


class RatingSchema(Schema):
    id = fields.String(description="UUID", dump_only=True)
    comment = fields.String(description="Comment")
    rating = fields.Integer(description="Rating")
    rating_date = fields.DateTime(description="Rating Date", dump_only=True)
    song_id = fields.String(description="Song UUID", dump_only=True)
    user_id = fields.String(description="User UUID")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("comment") or obj.get("comment") == "") and \
               (not obj.get("rating")) and \
               (not obj.get("rating_date")or obj.get("rating_date") == "") and \
               (not obj.get("song_id") or obj.get("song_id") == "") and \
               (not obj.get("user_id") or obj.get("user_id") == "")

# Schéma de base pour les évaluations
class BaseRatingSchema(Schema):
    comment = fields.String(description="Comment")
    rating = fields.Integer(description="Rating")

# Schéma pour la mise à jour des évaluations
class RatingUpdateSchema(BaseRatingSchema):
    # Permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("comment" in data and data["comment"] != "") or
                ("rating" in data)):
            raise ValidationError("at least one of ['comment', 'rating'] must be specified")
