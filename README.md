
# USTB-API
北科大教务相关 API 接口文档

 - Go 语言编写
 - 项目地址：[USTB-TeachSystem-API](https://github.com/jeasonstudio/USTB-TeachSystem-API)
 - 提BUG或新API需求，请去Github仓库提issue，或发送邮件到[me@jeasonstudio.cn](mailto:me@jeasonstudio.cn)
 - HOSTPATH: http://jeasonstudio.cn:8080
 - 文档：[点我](https://jeasonstudio.github.io/USTB-TeachSystem-API/)
 - Author: Jeason

目录：
 - [新版教务管理系统相关](README.md#L37)
   - [获取所有成绩](README.md#L22)
   - [获取创新学分](#L78)
   - [获取所有课程](#L119)
 - [开发相关]()


## 新版教务管理系统相关(/v1)

---

### @ 获取所有成绩(/classScore.ustbsu)

#### HOST: http://jeasonstudio.cn:8080/v1/classScore.ustbsu

### METHOD: GET

#### 请求REQUEST:

编号 | 参数名 | 值类型 | 是否必须 | 备注
---|---|---|---|---
1 | username | string | 是 | 学号 
2 | password | string | 是 | 新版教务管理系统密码

#### 响应RESPONSE: 

```json
{
    "avScore": "76.4",
    "body": [{
        "className": "基础外语I",
        "classNum": "10903021",
        "classType": "人社管必",
        "finalScore": "85",
        "firstScore": "85",
        "flag": "",
        "learnHour": "64",
        "semestre": "2015-2016-1 ",
        "stuScore": "4"
    }, {
        "className": "计算机算法设计",
        "classNum": "4248002",
        "classType": "本专业选",
        "finalScore": "91",
        "firstScore": "91",
        "flag": "",
        "learnHour": "16",
        "semestre": "2015-2016-1 ",
        "stuScore": "1"
    }],
    "gpa": "2.81"
}
```

编号 | 参数名 | 值类型 | 备注
---|---|---|---|---
1 | avScore | string | 加权平均分 
2 | gpa | string | GPA
3 | className | string | 课程名
4 | classNum | string | 课程号	
5 | classType | string | 课程类别	
6 | finalScore | string | 最终成绩(对外成绩单)		
7 | firstScore | string | 第一次成绩(排名用)	
8 | flag | string | 重修补考标志
9 | learnHour | string | 学时
10 | semestre | string | 学年学期
11 | stuScore | string | 学分

---

### @ 获取创新学分(/cxScore.ustbsu)

#### HOST: http://jeasonstudio.cn:8080/v1/cxScore.ustbsu

### METHOD: GET

#### 请求REQUEST:

编号 | 参数名 | 值类型 | 是否必须 | 备注
---|---|---|---|---
1 | username | string | 是 | 学号 
2 | password | string | 是 | 新版教务管理系统密码

#### 响应RESPONSE: 

```json
[{
    "cxType": "报告",
    "insertTime": "",
    "name": "APP开发与行业趋势",
    "score": "0.2",
    "semestre": " "
}, {
    "cxType": "报告",
    "insertTime": "",
    "name": "大数据时代与微软云计算",
    "score": "0.2",
    "semestre": " "
}]
```

编号 | 参数名 | 值类型 | 备注
---|---|---|---|---
1 | cxType | string | 创新学分类型
2 | insertTime | string | 录入时间
3 | name | string | 课程名
4 | score | string | 学分
5 | semestre | string | 学期学年

---

### @ 获取所有课程(/course.ustbsu)

#### HOST: http://jeasonstudio.cn:8080/v1/course.ustbsu

### METHOD: GET

#### 请求REQUEST:

编号 | 参数名 | 值类型 | 是否必须 | 备注
---|---|---|---|---
1 | username | string | 是 | 学号 
2 | password | string | 是 | 新版教务管理系统密码
3 | semestre | string | 是 | 学年学期，比如：2016-2017-2

#### 响应RESPONSE: 

> **提示：** 由于此 json 接口是原来的不知道那个老师写的，其中的一些 KEY 值为汉语拼音缩写，大家自己理解。

```json
{
    "selectedCourses": [{
        "SKRS": "97",
        "XS": "32",
        "DYXF": "2",
        "TJR": "41524122",
        "XH": "41524122",
        "SFQK": null,
        "SJCJ": null,
        "BZ": null,
        "XM": "赵吉彤",
        "JXBJ": "计1501",
        "XB": "男",
        "DYXS": "32",
        "SFYXTX": "1",
        "PTK": [],
        "XF": "2",
        "RXNJ": "2015",
        "BXJZRQ": "2016-09-15",
        "ID": 17613117,
        "KCLBM": "公共选修",
        "DYKCH": "1089096",
        "SKXSB_ID": 17984823,
        "KHLX": null,
        "ZYFX": "计算机科学与技术",
        "JTLB": "普通",
        "SKSJDDSTR": "(周4,第6节,1-16周 逸夫楼705) ",
        "QZ": 0,
        "SFYXCX": "1",
        "XXK": [],
        "SKSJDD": {
            "66": ["逸夫楼705", "1-16周"],
            "318": ["逸夫楼705", "1-16周"],
            "570": ["逸夫楼705", "1-16周"],
            "486": ["逸夫楼705", "1-16周"],
            "24": ["逸夫楼705", "1-16周"],
            "654": ["逸夫楼705", "1-16周"],
            "192": ["逸夫楼705", "1-16周"],
            "360": ["逸夫楼705", "1-16周"],
            "444": ["逸夫楼705", "1-16周"],
            "276": ["逸夫楼705", "1-16周"],
            "150": ["逸夫楼705", "1-16周"],
            "612": ["逸夫楼705", "1-16周"],
            "108": ["逸夫楼705", "1-16周"],
            "234": ["逸夫楼705", "1-16周"],
            "402": ["逸夫楼705", "1-16周"],
            "528": ["逸夫楼705", "1-16周"]
        },
        "DYKCM": "身边的法律",
        "KCCJ": null,
        "KXH": "1001",
        "SFKX": "1",
        "SSNJ": "2015",
        "KRL": 100,
        "SFXZRS": null,
        "KCLB": "7",
        "SFCX": null,
        "PSCJ": null,
        "XSXZ": "普通学生",
        "SFSEF": "0",
        "XNXQ": "2016-2017-1",
        "JSM": [{
            "JSM": "徐铭勋"
        }],
        "FREERS": "0",
        "JTZT": "10",
        "KCM": "身边的法律",
        "ZJT_ID": null,
        "TKJZRQ": "2016-09-19",
        "KCH": "1089096"
    }],
    "xnxqLessonPerDay": 6
}
```

## 开发相关

安装依赖：

```Go
go get github.com/astaxie/beego
go get github.com/PuerkitoBio/goquery
go get github.com/bitly/go-simplejson
```

开发调试：
```bash
bee run
```

打包发布：
```bash
bee pack -be GOOS=linux GOARCH=amd64
```

上传到服务器

执行，使用 Linux 做进程守护：
```bash
nohup ./USTB-TeachSystem-API &
```