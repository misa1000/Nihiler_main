﻿# Termbin开发日志
1.首先由于不了解命令行
``` Linux
echo "kinan" | curl -F c=@- localhost:888/create -s
而curl -F "c=kiana" localhost:888/create则确实是利用表单进行传参
```
具体的传输规则，在使用postform方法并没有成功从c中取得参数，在了解到该命令通过post请求体传输以后，使用bind方法解决了这个问题。

2. 其次在使用orm库进行数据库操作的时候，由于并不熟悉CRUD的细节语法操作，耽误了不少的时间进行测试寻找错误。

3.2. 在第二阶段，实现了登陆登出，访问属性的API设计。

进行API的设计，使用session记录用户的登陆信息从而使得用户在登陆时候所能使用一些特殊的功能。
---
而新的功能相对于原来基本相同但是有所差别，因此设计再开创一个专门的表用来记录用户登陆的时候所创建的内容，同时所有用户的内容放在同一个表里面，由此便于操作未被设置访问权限的相关数据。

3. 由于一些功能受到权限管辖，只有用户登录以后才可以进行操作，因此使用--中间件--来对请求进行拦截，同时检验session中是否有用户的登陆信息，通过这种手段开放登陆后才可以使用的有关功能。


### 关于登陆以后权限设置的实现
1. 只有该账户下的内容可以被该账户进行访问权限设置，而其他账户不可以对此进行访问的设置！！！

2. 进行对username的检验，从而判断是否有权限进行相关的操作。

3. **同时要注意到需要实现的两个功能其实是互相冲突的关系，设置了仅主人可见就不能再设置别人可见，而一旦设置了别人可见，就不可同时设置仅主人可见。** 同时我们应该注意到设置功能并不受到其他功能的影响，比如owner访问权限在another设置的时候仍然可以再设置，但是此时应该把another修改过来。

4.**登陆后用户可以对剪切板进行修改的几种情况:这是用户自己的剪切板。**
还有一种情况是查询，分为以下几种情况:
1.是作者本人的剪切板
2.不是作者本人但是作者没有设置访问权限的。
3.不是作者本人但是作者赋予了自己权限的。


### 对于第二阶段注册功能的测试：
1. 注册
```
curl -X POST localhost:888/user/register/WangZe/123456
```
2. 登陆
```
curl -X POST localhost:888/user/login/WangZe/123456
```
同时测试一下登陆不存在的用户

3. 登陆后创建剪切板
```
curl -X POST -F "c=暗区突围" localhost:888/WangZe/create
```
同时测试一下没有登陆的时候创建剪切板的后果

4. 再测试一下登陆后的删除，更新，查询功能
```
curl -X POST  -F "c=造梦西游4" localhost:888/WangZe/update/6
curl -X POST localhost:888/WangZe/query/6
curl -X POST localhost:888/WangZe/delete/6
```
5. 进行后两个功能的测试：对于访问权限的测试
```
设置仅作者可见：curl -X POST localhost:888/WangZe/owner/6/yes
设置仅作者和被设置的某个人可见：curl -X POST localhost:888/WangZe/another/6/rengengchen(同时测试一下没有权限的人查询会怎么样)
再使用被设置的人查询
```
在测试的时候根据代码解释没有权限，于是开始排查错误。
错误已经找到，原因是选中的实验数据已经被删除了，但是之前一直没有发现**=͟͟͞͞(꒪ᗜ꒪ ‧̣̥̇)**

### 第三阶段对阅读以后即时销毁功能的实现
1. 功能触发形式：表单c传参key：Misaka_Mikoto。
要求只有被赋予访问权限的人才可以使用
具体实现是在UserQuery中加入对用户名和key的检测
```
测试语句：curl -X POST -F "c=Misaka_Mikoto" localhost:888/query/:id
```

2. 决定设计一个新的创建功能函数，同时可以接受传参用来设置销毁时间。在这个处理函数里面加入延时操作，时间到以后直接从数据库里面删除这个剪切板！！！
```
测试用例：curl -F "c=暗区突围" -F "sunset=40" localhost:888/WangZe/create_Delay
40s以后再次测试
进行查询
```
在本次任务中首次使用了并发编程，为延时操作设置了一个goroutine来执行，从而达到，创建后返回信息，同时时间过了以后再次查询的时候无法查询的效果。

### 对于额外功能的实现
1. 实现项目的别名。这个问题的关键在于条件判断访问传参的类型，是id还是别名或者short。
```
测试用例：首先登陆
curl -X POST localhost:888/user/login/WangZe/123456
然后创建的时候加上别名：curl -X POST localhost:888/WangZe/create/mom
然后进行测试：curl -X POST localhost:888/WangZe/query/mom
```
