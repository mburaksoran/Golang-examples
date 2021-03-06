# Golang-examples

## 1. CRUD-API 

Example of API

### 1.1. TO-DO

- Version 1 - Fiber - GORM 
    - GET :heavy_check_mark: 
    - POST :heavy_check_mark:
    - PUT :heavy_check_mark:
    - DELETE :heavy_check_mark:
  
    
#### 1.1.1. Usage 

## End-point: CreateUser
### Method: POST
>```
>http://localhost:8080/api/createuser/
>```

*********************************************************************************************************************

## End-point: GetUser
### Method: GET
>```
>http://localhost:8080/api/getuser
>```

*********************************************************************************************************************

## End-point: GetUserbyid
### Method: GET
>```
>http://localhost:8080/api/getuser/:id
>```

*********************************************************************************************************************

## End-point: DeleteUser
### Method: DELETE
>```
>http://localhost:8080/api/deleteuser/:id
>```

*********************************************************************************************************************

## End-point: UpdateUser
### Method: PUT
>```
>http://localhost:8080/api/updateuser/:id
>```

*********************************************************************************************************************


## 2. HTTP Requests
- A HTTP request contains the following elements: 
 - Request line
 - Header fields
 - Message body(Optional)

### 2.1. HTTP Request Line
- Example of Request line **GET /software/htp/cics/index.html HTTP/1.1**
 - For this example the method is **GET**
 - the path is **/software/htp/cics/index.html**
 - HTTP version is **HTTP/1.1**

-Some of HTTP verbs

|       USE CASE        |  HTTP METHOD  |
|-----------------------|---------------|
| Create                | `POST`        |
| Read                  | `GET`         |
| Update                | `PUT`         |
| Partial Update        | `PATCH`       |
| Delete or Soft Delete | `DELETE`      |

### 2.2. HTTP Request Headers	
- Request-header fields allow the client to pass additional information about the request, and about the client itself, to the server.
 
| SOME OF HEADER FIELDS |
|-----------------------|
| [Accept-Charset](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Charset)        | 
| [Accept-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding)       |         
| [Accept-Language](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language)       |         
| [Authorization](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization)         | 
| [Expect](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect)                | 
| [From](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/From)                  | 
| [Host](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host)                  | 
| [If-Match](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match)              | 
| [If-Modified-Since](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Modified-Since)     | 
| [If-None-Match](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-None-Match)         | 
| [If-Range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Range)              | 
| [If-Unmodified-Since](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Unmodified-Since)   | 
| [Max-Forwards](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Max-Forwards)          | 
| [Proxy-Authorization](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authorization)   | 
| [Range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range)                | 
| [Referer](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer)              | 
| [TE](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/TE)              | 
| [User-Agent](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent)         | 

## 3. HTTP Response
- After receiving and interpreting a request message, a server responds with an HTTP response message. HTTP response include: 
 - Status line
 - Header fields 	
 - Message body(Optional)

### 3.1. Status Line
 - Example of Response Status line  ** HTTP/1.1 200 OK ** 
  - the HTTP version is HTTP/1.1
  - the status code is 200
  - the reason phrase is OK

#HTTP Response Status Codes 

- Informational (starting with numeric 1, 1xx)
- Success (starting with numeric 2, 2xx)
- Redirection (starting with numeric 3, 3xx)
- Client Error (starting with numeric 4, 4xx)
- Server Error (starting with numeric 5, 5xx)

### 3.1.1. 1xx: Information responses
| **HTTP STATUS CODE** |                                                                 |
|----------------------|-----------------------------------------------------------------|
| **100**              | Continue                                                        |
| **101**              | Switching Protocols                                             |
| **102**              | Processing                                                      |
| **103**              | Early Hints                                                     |

### 3.1.2. 2xx: Success status
| **HTTP STATUS CODE** |                                                                 |
|----------------------|-----------------------------------------------------------------|
| **200**              | OK                                                              |
| **201**              | Created                                                         |
| **202**              | Accepted                                                        |
| **203**              | NonAuthoritativeInfo                                            |
| **204**              | No Content                                                      |
| **205**              | Reset Content                                                   |
| **206**              | Partial Content                                                 |
| **207**              | Multi Status                                                    |
| **208**              | AlreadyReported                                                 |
| **226**              | IM Used                                                         |

### 3.1.3. 3xx: Redirection status
| **HTTP STATUS CODE** |                                                                 |
|----------------------|-----------------------------------------------------------------|
| **300**              | Multiple Choice                                                 |
| **301**              | Moved Permanently                                               |
| **302**              | Found                                                           |
| **303**              | See Other                                                       |
| **304**              | Not Modified                                                    |
| **305**              | Use Proxy                                                       |
| **306**              | Unused                                                          |
| **307**              | Temporary Redirect                                              |
| **308**              | Permanent Redirect                                              |

### 3.1.4. 4xx: Failure (User Input Related Failures)
| **HTTP STATUS CODE** |                                                                 |
|----------------------|-----------------------------------------------------------------|
| **400**              | Bad Request                                                     |
| **401**              | Unauthorized                                                    |
| **402**              | Payment Required                                                |
| **403**              | Forbidden                                                       |
| **404**              | Not Found                                                       |
| **405**              | Method Not Allowed                                              |
| **406**              | Not Acceptable                                                  |
| **408**              | Request Timeout                                                 |
| **409**              | Conflict                                                        |
| **410**              | Gone                                                            |
| **411**              | Length Required                                                 |
| **412**              | Precondition Failed                                             |
| **413**              | Payload Too Large                                               |
| **414**              | URI Too Long                                                    |
| **421**              | Misdirected Request                                             |
| **451**              | Unavailable For Legal Reasons                                   |

### 3.1.5. 5xx: Failure (System or Server Side Failures)
| **HTTP STATUS CODE** |                                                                 |
|----------------------|-----------------------------------------------------------------|
| **500**              |  Internal Server Error                                          |
| **501**              |  Not Implemented                         			     |
| **502**              |  Bad Gateway                              			     |
| **503**              |  Service Unavailable                     			     |
| **504**              |  Gateway Timeout     		                 			     |
| **505**              |  HTTP Version Not Supported 		                  	     |
| **506**              |  Variant Also Negotiates 		                  	     | 
| **507**              |  Insufficient Storage   		                  	     | 
| **508**              |  Loop Detected          		                  	     | 
| **510**              |  Not Extended          		                  	     | 
| **511**              |  Network Authentication Required	                  	     | 


### 3.2. Header Fields
- Header fields give information about the server and about further access to the resource identified by the Request-URI

| SOME OF HEADER FIELDS |  	
|-----------------------| 
| [Accept-Ranges](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges)		|
| [Age](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Age)			|
| [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag)			|
| [Location](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Location)			|
| [Proxy-Authenticate](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authenticate)	|
| [Retry-After](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After)		|
| [Server](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server)			|
| [Vary](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary)			|
| [WWW-Authenticate](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate) 	|
