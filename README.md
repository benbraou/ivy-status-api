# ivy-status-api
https://ivy-status-api.appspot.com/v1/status

## Introduction
`Ivy` is the new Angular renderer focused on speed and size reduction. It is published as an experimental API as of Angular 6.
As of April 30th 2018, Ivy is under active development by the Angular team.  
Ivy status is regularly updated by the Angular team in a dedicated [STATUS.md](https://github.com/angular/angular/blob/master/packages/core/src/render3/STATUS.md)  

The goal of this project is to parse this file and return the status as a JSON API.

## The project
An API written in go deployed using Google Cloud Platform.
It can be accessed through the URL: `https://ivy-status-api.appspot.com/v1/status`
### Response structure

```json
{
    "API": {
        "version": 1
    },
    "errors": null,
    "data": {
        "featureGroup": {
            "data": {
                "name": "root",
                "features": []
            },
            "featureGroups":[
                
            ]
        },
        "lastUpdateDate": "2018-04-29T12:21:22"
    }
}
```

```
// FeatureGroup can hold child feature groups. It is in reality a tree. e.g.
//                 [Implementation Status]
//                          / \
//                         /   \
//                        /     \
//                       /       \
//     [`@angular/compiler-cli`  [`@angular/compiler` changes]
//                changes]
//                  /
//                 /
//   [`ngtsc` TSC compiler transformer]
```
In the current documentation status, more information can be found in the models defined in the [source code](https://github.com/benbraou/ivy-status-api/blob/main/model/feature.go)  

Every feature group consists of one or several features. In the following example, the `Decorators` feature group consists of three features whose status is granular according to Runtime, Spec, Compiler and Back Patch

| Annotation          | `defineXXX()`                  | Run time | Spec     | Compiler | Back Patch |
| -------------------- | ------------------------------ | ------- | -------- | -------- | -------- |
| `@Component`         | ✅ `defineComponent()`         |    ✅    |  ✅      |  ✅      |  ❌      |
| `@Directive`         | ✅ `defineDirective()`         |    ✅    |  ✅      |  ✅      |  ❌      |
| `@Directive`         | ❌ `defineAbstractDirective()` |    ❌    |  ❌      |  ❌      |  ❌      |

The first feature in this example will be returned by the API in this form
```json

                  {
                    "name": "`@Component`",
                    "status": {
                      "completed": false,
                      "categories": [
                        "Annotation",
                        "`defineXXX()`",
                        "Run time",
                        "Spec",
                        "Compiler",
                        "Back Patch"
                      ],
                      "granularStatuses": [
                        {
                          "category": "`defineXXX()`",
                          "code": "IMPLEMENTED",
                          "description": "`defineComponent()`"
                        },
                        {
                          "category": "Run time",
                          "code": "IMPLEMENTED",
                          "description": "Run time"
                        },
                        {
                          "category": "Spec",
                          "code": "IMPLEMENTED",
                          "description": "Spec"
                        },
                        {
                          "category": "Compiler",
                          "code": "IMPLEMENTED",
                          "description": "Compiler"
                        },
                        {
                          "category": "Back Patch",
                          "code": "NOT_IMPLEMENTED",
                          "description": "Back Patch"
                        }
                      ]
                    },
                    "childFeatures": null}
                 
```

## Contribution

Please feel free to open [an issue](https://github.com/benbraou/ivy-status-api/issues?state=open).  
Pull requests with the fix and a test are welcome.  
To launch tests: `scripts.test.sh`  



