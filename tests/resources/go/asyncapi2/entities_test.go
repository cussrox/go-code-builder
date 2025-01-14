package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/yudai/gojsondiff/formatter"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := Info{
		Version: "v1",
		MapOfAnything: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := Info{
		Version: "v1",
	}

	res, err := json.Marshal(i)
	require.NoError(t, err)
	assert.Equal(t, `{"version":"v1"}`, string(res))
}

func TestAsyncAPI_MarshalJSON(t *testing.T) {
	data := `{
    "asyncapi": "2.0.0",
    "info": {
        "title": "Streetlights API",
        "version": "1.0.0",
        "description": "The Smartylighting Streetlights API allows you to remotely manage the city lights.\n\n### Check out its awesome features:\n\n* Turn a specific streetlight on/off \ud83c\udf03\n* Dim a specific streetlight \ud83d\ude0e\n* Receive real-time information about environmental lighting conditions \ud83d\udcc8\n",
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0"
        }
    },
    "servers": {
        "production": {
            "url": "test.mosquitto.org:{port}",
            "protocol": "mqtt",
            "description": "Test broker",
            "variables": {
                "port": {
                    "description": "Secure connection (TLS) is available through port 8883.",
                    "default": "1883",
                    "enum": [
                        "1883",
                        "8883"
                    ]
                }
            },
            "security": [
                {
                    "apiKey": []
                },
                {
                    "supportedOauthFlows": [
                        "streetlights:on",
                        "streetlights:off",
                        "streetlights:dim"
                    ]
                },
                {
                    "openIdConnectWellKnown": []
                }
            ]
        }
    },
    "defaultContentType": "application/json",
    "channels": {
        "smartylighting/streetlights/1/0/event/{streetlightId}/lighting/measured": {
            "description": "The topic on which measured values may be produced and consumed.",
            "parameters": {
                "streetlightId": {
                    "$ref": "#/components/parameters/streetlightId"
                }
            },
            "publish": {
                "summary": "Inform about environmental lighting conditions of a particular streetlight.",
                "operationId": "receiveLightMeasurement",
                "traits": [
                    {
                        "$ref": "#/components/operationTraits/kafka"
                    }
                ],
                "message": {
                    "$ref": "#/components/messages/lightMeasured"
                }
            }
        },
        "smartylighting/streetlights/1/0/action/{streetlightId}/turn/on": {
            "parameters": {
                "streetlightId": {
                    "$ref": "#/components/parameters/streetlightId"
                }
            },
            "subscribe": {
                "operationId": "turnOn",
                "traits": [
                    {
                        "$ref": "#/components/operationTraits/kafka"
                    }
                ],
                "message": {
                    "$ref": "#/components/messages/turnOnOff"
                }
            }
        },
        "smartylighting/streetlights/1/0/action/{streetlightId}/turn/off": {
            "parameters": {
                "streetlightId": {
                    "$ref": "#/components/parameters/streetlightId"
                }
            },
            "subscribe": {
                "operationId": "turnOff",
                "traits": [
                    {
                        "$ref": "#/components/operationTraits/kafka"
                    }
                ],
                "message": {
                    "$ref": "#/components/messages/turnOnOff"
                }
            }
        },
        "smartylighting/streetlights/1/0/action/{streetlightId}/dim": {
            "parameters": {
                "streetlightId": {
                    "$ref": "#/components/parameters/streetlightId"
                }
            },
            "subscribe": {
                "operationId": "dimLight",
                "traits": [
                    {
                        "$ref": "#/components/operationTraits/kafka"
                    }
                ],
                "message": {
                    "$ref": "#/components/messages/dimLight"
                }
            }
        }
    },
    "components": {
        "messages": {
            "lightMeasured": {
                "name": "lightMeasured",
                "title": "Light measured",
                "summary": "Inform about environmental lighting conditions of a particular streetlight.",
                "contentType": "application/json",
                "traits": [
                    {
                        "$ref": "#/components/messageTraits/commonHeaders"
                    }
                ],
                "payload": {
                    "$ref": "#/components/schemas/lightMeasuredPayload"
                }
            },
            "turnOnOff": {
                "name": "turnOnOff",
                "title": "Turn on/off",
                "summary": "Command a particular streetlight to turn the lights on or off.",
                "traits": [
                    {
                        "$ref": "#/components/messageTraits/commonHeaders"
                    }
                ],
                "payload": {
                    "$ref": "#/components/schemas/turnOnOffPayload"
                }
            },
            "dimLight": {
                "name": "dimLight",
                "title": "Dim light",
                "summary": "Command a particular streetlight to dim the lights.",
                "traits": [
                    {
                        "$ref": "#/components/messageTraits/commonHeaders"
                    }
                ],
                "payload": {
                    "$ref": "#/components/schemas/dimLightPayload"
                }
            }
        },
        "schemas": {
            "lightMeasuredPayload": {
                "type": "object",
                "properties": {
                    "lumens": {
                        "type": "integer",
                        "minimum": 0,
                        "description": "Light intensity measured in lumens."
                    },
                    "sentAt": {
                        "$ref": "#/components/schemas/sentAt"
                    }
                }
            },
            "turnOnOffPayload": {
                "type": "object",
                "properties": {
                    "command": {
                        "type": "string",
                        "enum": [
                            "on",
                            "off"
                        ],
                        "description": "Whether to turn on or off the light."
                    },
                    "sentAt": {
                        "$ref": "#/components/schemas/sentAt"
                    }
                }
            },
            "dimLightPayload": {
                "type": "object",
                "properties": {
                    "percentage": {
                        "type": "integer",
                        "description": "Percentage to which the light should be dimmed to.",
                        "minimum": 0,
                        "maximum": 100
                    },
                    "sentAt": {
                        "$ref": "#/components/schemas/sentAt"
                    }
                }
            },
            "sentAt": {
                "type": "string",
                "format": "date-time",
                "description": "Date and time when the message was sent."
            }
        },
        "securitySchemes": {
            "apiKey": {
                "type": "apiKey",
                "in": "user",
                "description": "Provide your API key as the user and leave the password empty."
            },
            "supportedOauthFlows": {
                "type": "oauth2",
                "description": "Flows to support OAuth 2.0",
                "flows": {
                    "implicit": {
                        "authorizationUrl": "https://authserver.example/auth",
                        "scopes": {
                            "streetlights:on": "Ability to switch lights on",
                            "streetlights:off": "Ability to switch lights off",
                            "streetlights:dim": "Ability to dim the lights"
                        }
                    },
                    "password": {
                        "tokenUrl": "https://authserver.example/token",
                        "scopes": {
                            "streetlights:on": "Ability to switch lights on",
                            "streetlights:off": "Ability to switch lights off",
                            "streetlights:dim": "Ability to dim the lights"
                        }
                    },
                    "clientCredentials": {
                        "tokenUrl": "https://authserver.example/token",
                        "scopes": {
                            "streetlights:on": "Ability to switch lights on",
                            "streetlights:off": "Ability to switch lights off",
                            "streetlights:dim": "Ability to dim the lights"
                        }
                    },
                    "authorizationCode": {
                        "authorizationUrl": "https://authserver.example/auth",
                        "tokenUrl": "https://authserver.example/token",
                        "refreshUrl": "https://authserver.example/refresh",
                        "scopes": {
                            "streetlights:on": "Ability to switch lights on",
                            "streetlights:off": "Ability to switch lights off",
                            "streetlights:dim": "Ability to dim the lights"
                        }
                    }
                }
            },
            "openIdConnectWellKnown": {
                "type": "openIdConnect",
                "openIdConnectUrl": "https://authserver.example/.well-known"
            }
        },
        "parameters": {
            "streetlightId": {
                "description": "The ID of the streetlight.",
                "schema": {
                    "type": "string"
                }
            }
        },
        "messageTraits": {
            "commonHeaders": {
                "headers": {
                    "type": "object",
                    "properties": {
                        "my-app-header": {
                            "type": "integer",
                            "minimum": 0,
                            "maximum": 100
                        }
                    }
                }
            }
        },
        "operationTraits": {
            "kafka": {
                "bindings": {
                    "kafka": {
                        "clientId": "my-app-id"
                    }
                }
            }
        }
    }
}`

	var a AsyncAPI
	require.NoError(t, json.Unmarshal([]byte(data), &a))

	marshaled, err := json.Marshal(a)
	require.NoError(t, err)
	aj := assertjson.Comparer{
		FormatterConfig: formatter.AsciiFormatterConfig{
			Coloring: true,
		},
	}
	aj.Equal(t, []byte(data), marshaled)
}
