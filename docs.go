
package main
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "apis": [
        {
            "path": "/urls",
            "description": "URL management API"
        }
    ],
    "info": {
        "title": "WUS",
        "description": "WID URL shortener"
    }
}`
var apiDescriptionsJson = map[string]string{"urls":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://127.0.0.1:3000",
    "resourcePath": "/urls",
    "apis": [
        {
            "path": "/urls/{some_id}",
            "description": "Get URL by ID",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "FindByID",
                    "type": "string",
                    "items": {},
                    "summary": "Get URL by ID",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "some_id",
                            "description": "Some ID",
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
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.solher.zest.apierrors.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.solher.zest.apierrors.APIError": {
            "id": "github.com.solher.zest.apierrors.APIError",
            "properties": {
                "description": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "errorCode": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "raw": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "status": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
