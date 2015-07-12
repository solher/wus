
package main
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "apis": [
        {
            "path": "/urls",
            "description": "Url resource"
        },
        {
            "path": "/accounts",
            "description": "Account resource"
        },
        {
            "path": "/aclMappings",
            "description": "AclMapping resource"
        },
        {
            "path": "/users",
            "description": "User resource"
        },
        {
            "path": "/roleMappings",
            "description": "RoleMapping resource"
        },
        {
            "path": "/sessions",
            "description": "Session resource"
        },
        {
            "path": "/roles",
            "description": "Role resource"
        },
        {
            "path": "/acls",
            "description": "Acl resource"
        }
    ],
    "info": {
        "title": "WUS",
        "description": "WID URL Shortener"
    }
}`
var apiDescriptionsJson = map[string]string{"aclMappings":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/aclMappings",
    "apis": [
        {
            "path": "/aclMappings",
            "description": "Create one or multiple AclMapping instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple AclMapping instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.AclMapping",
                    "items": {},
                    "summary": "Find all AclMapping instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple AclMapping instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all AclMapping instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/aclMappings/{id}",
            "description": "Find a AclMapping instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.AclMapping",
                    "items": {},
                    "summary": "Find a AclMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a AclMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a AclMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/aclMappings/{pk}/{relatedResource}",
            "description": "Create one or multiple AclMapping instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple AclMapping instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.AclMapping",
                    "items": {},
                    "summary": "Find all AclMapping instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple AclMapping instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all AclMapping instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/aclMappings/{pk}/{relatedResource}/{fk}",
            "description": "Find a AclMapping instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.AclMapping",
                    "items": {},
                    "summary": "Find a AclMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a AclMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "AclMapping",
                            "description": "AclMapping instance data",
                            "dataType": "github.com.solher.zest.domain.AclMapping",
                            "type": "github.com.solher.zest.domain.AclMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.AclMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a AclMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "AclMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"users":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/users",
    "apis": [
        {
            "path": "/users",
            "description": "Create one or multiple User instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple User instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance(s) data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.User",
                    "items": {},
                    "summary": "Find all User instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple User instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance(s) data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all User instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/users/{id}",
            "description": "Find a User instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.User",
                    "items": {},
                    "summary": "Find a User instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a User instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a User instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/users/{pk}/{relatedResource}",
            "description": "Create one or multiple User instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple User instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance(s) data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.User",
                    "items": {},
                    "summary": "Find all User instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple User instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance(s) data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all User instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/users/{pk}/{relatedResource}/{fk}",
            "description": "Find a User instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.User",
                    "items": {},
                    "summary": "Find a User instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a User instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.User"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a User instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "User id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"roleMappings":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/roleMappings",
    "apis": [
        {
            "path": "/roleMappings",
            "description": "Create one or multiple RoleMapping instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple RoleMapping instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.RoleMapping",
                    "items": {},
                    "summary": "Find all RoleMapping instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple RoleMapping instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all RoleMapping instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roleMappings/{id}",
            "description": "Find a RoleMapping instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.RoleMapping",
                    "items": {},
                    "summary": "Find a RoleMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a RoleMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a RoleMapping instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roleMappings/{pk}/{relatedResource}",
            "description": "Create one or multiple RoleMapping instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple RoleMapping instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.RoleMapping",
                    "items": {},
                    "summary": "Find all RoleMapping instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple RoleMapping instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance(s) data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all RoleMapping instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roleMappings/{pk}/{relatedResource}/{fk}",
            "description": "Find a RoleMapping instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.RoleMapping",
                    "items": {},
                    "summary": "Find a RoleMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a RoleMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "RoleMapping",
                            "description": "RoleMapping instance data",
                            "dataType": "github.com.solher.zest.domain.RoleMapping",
                            "type": "github.com.solher.zest.domain.RoleMapping",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.RoleMapping"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a RoleMapping instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "RoleMapping id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"sessions":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/sessions",
    "apis": [
        {
            "path": "/sessions",
            "description": "Create one or multiple Session instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Session instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.Session",
                    "items": {},
                    "summary": "Find all Session instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Session instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Session instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/sessions/{id}",
            "description": "Find a Session instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.Session",
                    "items": {},
                    "summary": "Find a Session instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Session instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Session instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/sessions/{pk}/{relatedResource}",
            "description": "Create one or multiple Session instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Session instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.Session",
                    "items": {},
                    "summary": "Find all Session instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Session instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Session instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/sessions/{pk}/{relatedResource}/{fk}",
            "description": "Find a Session instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.Session",
                    "items": {},
                    "summary": "Find a Session instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Session instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Session",
                            "description": "Session instance data",
                            "dataType": "github.com.solher.zest.domain.Session",
                            "type": "github.com.solher.zest.domain.Session",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Session instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Session id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"roles":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/roles",
    "apis": [
        {
            "path": "/roles",
            "description": "Create one or multiple Role instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Role instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.Role",
                    "items": {},
                    "summary": "Find all Role instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Role instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Role instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roles/{id}",
            "description": "Find a Role instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.Role",
                    "items": {},
                    "summary": "Find a Role instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Role instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Role instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roles/{pk}/{relatedResource}",
            "description": "Create one or multiple Role instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Role instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.Role",
                    "items": {},
                    "summary": "Find all Role instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Role instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Role instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/roles/{pk}/{relatedResource}/{fk}",
            "description": "Find a Role instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.Role",
                    "items": {},
                    "summary": "Find a Role instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Role instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Role",
                            "description": "Role instance data",
                            "dataType": "github.com.solher.zest.domain.Role",
                            "type": "github.com.solher.zest.domain.Role",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Role"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Role instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Role id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"acls":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/acls",
    "apis": [
        {
            "path": "/acls",
            "description": "Create one or multiple Acl instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Acl instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.Acl",
                    "items": {},
                    "summary": "Find all Acl instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Acl instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Acl instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/acls/{id}",
            "description": "Find a Acl instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.Acl",
                    "items": {},
                    "summary": "Find a Acl instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Acl instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Acl instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/acls/{pk}/{relatedResource}",
            "description": "Create one or multiple Acl instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Acl instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.Acl",
                    "items": {},
                    "summary": "Find all Acl instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Acl instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Acl instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/acls/{pk}/{relatedResource}/{fk}",
            "description": "Find a Acl instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.Acl",
                    "items": {},
                    "summary": "Find a Acl instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Acl instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Acl",
                            "description": "Acl instance data",
                            "dataType": "github.com.solher.zest.domain.Acl",
                            "type": "github.com.solher.zest.domain.Acl",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Acl"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Acl instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Acl id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"urls":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/urls",
    "apis": [
        {
            "path": "/{shortHandle}",
            "description": "Redirect a short Url to the original one",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "RedirectUrl",
                    "type": "",
                    "items": {},
                    "summary": "Redirect a short Url to the original one",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "shortHandle",
                            "description": "Shortened Url",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 301,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/urls",
            "description": "Create one or multiple Url instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Url instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance(s) data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.wus.domain.Url",
                    "items": {},
                    "summary": "Find all Url instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Url instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance(s) data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Url instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/urls/{id}",
            "description": "Find a Url instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.wus.domain.Url",
                    "items": {},
                    "summary": "Find a Url instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Url instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Url instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/urls/{pk}/{relatedResource}",
            "description": "Create one or multiple Url instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Url instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance(s) data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.wus.domain.Url",
                    "items": {},
                    "summary": "Find all Url instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Url instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance(s) data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Url instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/urls/{pk}/{relatedResource}/{fk}",
            "description": "Find a Url instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.wus.domain.Url",
                    "items": {},
                    "summary": "Find a Url instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Url instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Url",
                            "description": "Url instance data",
                            "dataType": "github.com.solher.wus.domain.Url",
                            "type": "github.com.solher.wus.domain.Url",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.wus.domain.Url"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Url instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Url id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.wus.domain.Url": {
            "id": "github.com.solher.wus.domain.Url",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "enabled": {
                    "type": "bool",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "longUrl": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "notes": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "shortHandle": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,"accounts":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "localhost:3005/api",
    "resourcePath": "/accounts",
    "apis": [
        {
            "path": "/accounts/signin",
            "description": "Signin with an email and a password",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Signin",
                    "type": "github.com.solher.zest.domain.Session",
                    "items": {},
                    "summary": "Signin with an email and a password",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Credentials",
                            "description": "The user credentials",
                            "dataType": "github.com.solher.zest.resources.Credentials",
                            "type": "github.com.solher.zest.resources.Credentials",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Session"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts/signout",
            "description": "Signout an account",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Signout",
                    "type": "",
                    "items": {},
                    "summary": "Signout an account",
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts/signup",
            "description": "Signup an account",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Signup",
                    "type": "github.com.solher.zest.domain.Account",
                    "items": {},
                    "summary": "Signup an account",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "User",
                            "description": "User instance data",
                            "dataType": "github.com.solher.zest.domain.User",
                            "type": "github.com.solher.zest.domain.User",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts",
            "description": "Create one or multiple Account instances",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Account instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "Find",
                    "type": "github.com.solher.zest.domain.Account",
                    "items": {},
                    "summary": "Find all Account instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "Upsert",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Account instances",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAll",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Account instances matched by filter",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts/{id}",
            "description": "Find a Account instance",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "github.com.solher.zest.domain.Account",
                    "items": {},
                    "summary": "Find a Account instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByID",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Account instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByID",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Account instance",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts/{pk}/{relatedResource}",
            "description": "Create one or multiple Account instances of a related resource",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateRelated",
                    "type": "",
                    "items": {},
                    "summary": "Create one or multiple Account instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "GET",
                    "nickname": "FindRelated",
                    "type": "github.com.solher.zest.domain.Account",
                    "items": {},
                    "summary": "Find all Account instances  of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpsertRelated",
                    "type": "",
                    "items": {},
                    "summary": "Upsert one or multiple Account instances of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance(s) data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteAllRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete all Account instances of a related resource matched by filter",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/accounts/{pk}/{relatedResource}/{fk}",
            "description": "Find a Account instance of a related resource",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByIDRelated",
                    "type": "github.com.solher.zest.domain.Account",
                    "items": {},
                    "summary": "Find a Account instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "filter",
                            "description": "JSON filter defining fields and includes",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": false,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "PUT",
                    "nickname": "UpdateByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Update attributes of a Account instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "Account",
                            "description": "Account instance data",
                            "dataType": "github.com.solher.zest.domain.Account",
                            "type": "github.com.solher.zest.domain.Account",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.domain.Account"
                        }
                    ]
                },
                {
                    "httpMethod": "DELETE",
                    "nickname": "DeleteByIDRelated",
                    "type": "",
                    "items": {},
                    "summary": "Delete a Account instance of a related resource",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "pk",
                            "description": "Account id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "relatedResource",
                            "description": "Related resource name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "fk",
                            "description": "Related resource id",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 204,
                            "message": "Request was successful",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.domain.Account": {
            "id": "github.com.solher.zest.domain.Account",
            "properties": {
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.RoleMapping"
                    },
                    "format": ""
                },
                "sessions": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.Session"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "users": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.User"
                    },
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Acl": {
            "id": "github.com.solher.zest.domain.Acl",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "method": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "resource": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.AclMapping": {
            "id": "github.com.solher.zest.domain.AclMapping",
            "properties": {
                "acl": {
                    "type": "github.com.solher.zest.domain.Acl",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "aclId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Role": {
            "id": "github.com.solher.zest.domain.Role",
            "properties": {
                "aclMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.solher.zest.domain.AclMapping"
                    },
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleMappings": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "RoleMapping"
                    },
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.RoleMapping": {
            "id": "github.com.solher.zest.domain.RoleMapping",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "role": {
                    "type": "github.com.solher.zest.domain.Role",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "roleId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.Session": {
            "id": "github.com.solher.zest.domain.Session",
            "properties": {
                "account": {
                    "type": "github.com.solher.zest.domain.Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "agent": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "authToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "deletedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ip": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "validTo": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.domain.User": {
            "id": "github.com.solher.zest.domain.User",
            "properties": {
                "account": {
                    "type": "Account",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "accountId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "createdAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "firstName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "lastName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updatedAt": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.solher.zest.resources.Credentials": {
            "id": "github.com.solher.zest.resources.Credentials",
            "properties": {
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "password": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "rememberMe": {
                    "type": "bool",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
