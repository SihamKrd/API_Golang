from marshmallow import Schema, fields, validates_schema, ValidationError

# Schéma utilisateur de création
class UserRegistrationSchema(Schema):
    name = fields.String(description="Name")
    username = fields.String(description="Username")
    email = fields.Email(description="Email of the user")
    password = fields.String(description="Password")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("password") or obj.get("password") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("username") or obj.get("username") == "") and \
               (not obj.get("email") or obj.get("email") == "")
               

# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.String(description="UUID")
    name = fields.String(description="Name")
    username = fields.String(description="Username")
    email = fields.Email(description="Email of the user")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("username") or obj.get("username") == "") and \
               (not obj.get("email") or obj.get("email") == "")


class BaseUserSchema(Schema):
    name = fields.String(description="Name")
    password = fields.String(description="Password")
    username = fields.String(description="Username")
    email = fields.Email(description="Email of the user")


# Schéma utilisateur de modification (name, username, password)
class UserUpdateSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or
                ("username" in data and data["username"] != "") or
                ("password" in data and data["password"] != "") or
                ("email" in data and data["email"] != "")) :
            raise ValidationError("at least one of ['name','username','password','email'] must be specified")
