wxTokenCenter
===

[![Build Status](https://travis-ci.org/coderguang/wxTokenCenter.svg?branch=master)](https://travis-ci.org/coderguang/wxTokenCenter)
![](https://img.shields.io/badge/language-golang-orange.svg)
[![codebeat badge](https://codebeat.co/badges/50602d7c-235d-4353-bf0d-0f5933c6e121)](https://codebeat.co/projects/github-com-coderguang-wxtokencenter-master)
[![](https://img.shields.io/badge/wp-@royalchen-blue.svg)](https://www.royalchen.com)

## wx token manager tool
 * easy to manager more than one wx token
 * auto update token when it out of time
 * only need config your appid and secret
 * require by get or post
 

## how to star
### 1. clone repository 
```shell
git clone git@github.com:coderguang/wxTokenCenter.git
```

### 2. config your appid and secret in config/config.json
```json
{
    "port":"2100", //listen port
    "configs":[
    {
         "category":"categoryOne",  //require params
         "type":"typeOne", //require params
         "appid":"your_appid1",
         "secret":"your_secret1"
    },
   {
         "category":"categoryOne",
         "type":"typeTwo",
         "appid":"your_appid2",
          "secret":"your_secret2"
    },
    {
          "category":"categoryTwo",
          "type":"twoTypeOne",
          "appid":"your_appid3",
           "secret":"your_secret3"
     }
    ]
}
```

### 3. run the program
```shell
    go run main.go config/config.json
```
   if success,you will get output like below:
   
   ![init](https://github.com/coderguang/img/blob/master/wxTokenCenter/init.png)
    
### 4. client requier format
```shell
   curl -i https://your_ips/?key=gzh,yaohao
```
   you will get like below in client:
  ![require](https://github.com/coderguang/img/blob/master/wxTokenCenter/require.png)
  
   also,server will log this requires,like below:
   ![require_ok](https://github.com/coderguang/img/blob/master/wxTokenCenter/require_ok.png)
  
   ## About me

**Author** | _[royalchen](https://www.royalchen.com)_
---------- | -----------------
email  | royalchen@royalchen.com
qq  | royalchen@royalchen.com
website | [www.royalchen.com](https://www.royalchen.com)
