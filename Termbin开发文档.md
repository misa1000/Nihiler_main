# Termbin开发文档
# Termbin简介
是一个使用命令行实现剪切板功能的小工具，它在具有基本的增删查改的功能的基础上还支持登陆登出，访问属性限制，以及阅后即焚或者定时删除的功能。

---
### 具体实现
#### CRUD的最基本功能
 最基本的CRUD功能，需要设置相关的路由同时建立有关的连接数据库。保留记录信息，具体内容的哈希加密方式，是整个项目的核心基础功能，其他功能都是建立在这些之上而实现的。

使用指令为：
在linux命令行里面末尾加上-s防止出现很多其他信息以保证整洁
```基本CRUD
创建内容：curl -F "c=kiana" localhost:888/create -s
删除内容：curl -X DELETE localhost:888/:id -s
更新内容：curl -F "c=_Elysia_"  -X PUT localhost:888/:id -s
查询内容：curl localhost:888/:id -s
```


创建
![输入图片说明](/imgs/2024-03-29/4pB6AEeZHTtBuzDE.png)
更新(少截了两个字母ed)
![输入图片说明](/imgs/2024-03-29/kmzXVjnQxNpgPgJo.png)
查询
![输入图片说明](/imgs/2024-03-29/4r2kGsA5HC4ETOMi.png)
删除
![输入图片说明](/imgs/2024-03-29/rOgTJCP2GXLRPJyP.png)


#### 加入登录功能以后的CRUD功能以及访问权限的设置
这个阶段最明显的改观就是加入了对于访问权限的控制，用户可以注册并登录（当然注册登录过程非常简单）。

```注册登录
注册：curl -X POST localhost:888/user/register/MisakaMikoto/040502 -s
登录：curl -X POST localhost:888/user/login/MisakaMikoto/040502 -s
```
(为了防止用户密码泄露采用了哈希加密算法在数据库里面保存您的密码，因此无须担忧密码泄露)
注册
![输入图片说明](/imgs/2024-03-29/trYL1R0GhbwEPxlg.png)
登录
![输入图片说明](/imgs/2024-03-29/X6ddapp5pwUxkzzA.png)




##### 接下来是用户登录以后的CRUD操作，这些功能只有用户登录以后才可以使用，如果在未登录状态下使用用户进行操作，则会拒绝您的请求。（如果用户未注册就登录则会显示无此人。）

```登录以后的CRUD
创建：curl -X POST -F "c=Railgun" localhost:888/MisakaMikoto/create -s（增加了记录作者信息功能）
还有另一种创建方式（增加别名创建）：curl -X POST -F "c=Railgun" localhost:888/MisakaMikoto/create/misaka -s
```
![输入图片说明](/imgs/2024-03-29/kLVlJhQxz40mpSO2.png)
```
更新：curl -X POST -F "c=Railguns" localhost:888/MisakaMikoto/update/1 -s
```
![输入图片说明](/imgs/2024-03-29/JPDi1AfGDyelaM3f.png)
```
查询：可以使用多种方式进行查询，例如常用的id，或者short（url末端四字随机字符串），还有别名
curl -X POST localhost:888/MisakaMikoto/query/1 -s
curl -X POST localhost:888/MisakaMikoto/query/misaka -s
curl -X POST localhost:888/MisakaMikoto/query/qiap -s
```
![输入图片说明](/imgs/2024-03-29/afFaW8zDmH33G636.png)
```
删除:curl -X POST localhost:888/MisakaMikoto/delete/1 -s
```
![输入图片说明](/imgs/2024-03-29/qnOqCqe2q9YYrIO3.png)

#### 登录最重要的功能：对于访问权限的控制
1. 用户可以设置某个剪切板内容仅自己可以访问，那么其他人登录后也无法访问你的这个权限剪切板。(请注意一定设置为yes，如果不想要开启此功能则为no)
```
设置：curl -X POST localhost:888/MisakaMikoto/owner/2/yes -s
由于刚才的记录以及被删除，所以又创建了一个剪切板
```
![输入图片说明](/imgs/2024-03-29/bKk3mHHYPsxAQnBD.png)
接下来你可以尝试注册一个用户，使用这个用户对此进行查询，请求会被拒绝

2. 但是很多时候我们可能不想让它公开，但是想要给特定的朋友看，自己同时也能看（这是独属于你俩的秘密（很抱歉目前只支持让另一人看的功能**╮ (˃̶͈◡˂̶͈)）**）
```
设置：curl -X POST localhost:888/MisakaMikoto/another/2/MisakaMikot -s
```
此功能与前一功能实际冲突，因此不能够同时设置（设置一个另一个会自动修正，因此使用时候无须注意）

#### 一些额外功能的实现
1. 阅后即焚
顾名思义，就是查询以后立刻删除的功能。该功能针对设置另一人可看的功能设置，当另一个人查询立刻销毁。
该功能由被设置的人查询的时候选择触发（为什么没有让创建者控制触发呢？其实是因为没有原因**╮ (˃̶͈◡˂̶͈)）**）
```
设置：curl -F "c=Misaka_Mikoto" localhost:888/MisakaMikot/query/2 -s
之后再请求就会被拒绝
```

2. 超时销毁功能
用户在创建剪切板的时候可以设置一个时间（比如300s），在创建成功以后的300s以内是可以被查询到的，但是一旦超过该时间该剪切板就会自动销毁。
```
时间只支持秒数设置
设置销毁时间：curl -F "c=呱太" -F "sunset=60" localhost:888/MisakaMikoto/create_Delay -s
```


完结......个人感觉这个东西很开门，但是没什么用**╮(╯▽╰)╭**如果你觉得还不错，请给我一个免费的star！！！感激不尽！！！
