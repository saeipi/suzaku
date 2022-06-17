**一、PO :(persistant object )，持久对象**

可以看成是与数据库中的表相映射的java对象。使用Hibernate来生成PO是不错的选择。  
  
**二、VO :(value object) ，值对象**

通常用于业务层之间的数据传递，和PO一样也是仅仅包含数据而已。但应是抽象出的业务对象,可以和表对应,也可以不对应,这根据业务的需要。  
  
PO只能用在数据层，VO用在商务逻辑层和表示层。各层操作属于该层自己的数据对象，这样就可以降低各层之间的耦合，便于以后系统的维护和扩展。  
  
**三、DAO :(Data Access Objects)  ，数据访问对象接口  
**DAO是Data Access Object数据访问接口，数据访问：顾名思义就是与数据库打交道。夹在业务逻辑与数据库资源中间。  
  
J2EE开发人员使用数据访问对象（DAO）设计模式把底层的数据访问逻辑和高层的商务逻辑分开.实现DAO模式能够更加专注于编写数据访问代码。  
  
DAO模式是标准的J2EE设计模式之一.开发人员使用这个模式把底层的数据访问操作和上层的商务逻辑分开。一个典型的DAO实现有下列几个组件：  
  1. 一个DAO工厂类；  
  2. 一个DAO接口；  
  3. 一个实现DAO接口的具体类；  
  4. 数据传递对象（有些时候叫做值对象）。  
  具体的DAO类包含了从特定的数据源访问数据的逻辑。  
  
**四、BO :(Business Object)，业务对象层  
**表示应用程序领域内“事物”的所有实体类。这些实体类驻留在服务器上，并利用服务类来协助完成它们的职责。  
  
**五、POJO :(Plain Old Java Objects)，简单的Java对象  
**实际就是普通JavaBeans,使用POJO名称是为了避免和EJB混淆起来, 而且简称比较直接。  
其中有一些属性及其getter、setter方法的类,有时可以作为value object或dto(Data Transform Object)来使用.当然,如果你有一个简单的运算属性也是可以的,但不允许有业务方法,也不能携带有connection之类的方法。

- VO（View Object）：视图对象，用于展示层，它的作用是把某个指定页面（或组件）的所有数据封装起来。
- DTO（Data Transfer Object）：数据传输对象，这个概念来源于J2EE的设计模式，原来的目的是为了EJB的分布式应用提供粗粒度的数据实体，以减少分布式调用的次数，从而提高分布式调用的性能和降低网络负载，但在这里，我泛指用于展示层与服务层之间的数据传输对象。
- DO（Domain Object）：领域对象，就是从现实世界中抽象出来的有形或无形的业务实体。
- PO（Persistent Object）：持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么，数据表中的每个字段（或若干个）就对应PO的一个（或若干个）属性。

下面以一个时序图建立简单模型来描述上述对象在三层架构应用中的位置  
用户发出请求（可能是填写表单），表单的数据在展示层被匹配为VO。  
展示层把VO转换为服务层对应方法所要求的DTO，传送给服务层。  
服务层首先根据DTO的数据构造（或重建）一个DO，调用DO的业务方法完成具体业务。  
服务层把DO转换为持久层对应的PO（可以使用ORM工具，也可以不用），调用持久层的持久化方法，把PO传递给它，完成持久化操作。  
  
VO: value object, view object  
PO: persistent object  
QO: query object  
DAO: data access object-there is also DAO mode  
DTO: data transfer object-there is also DTO mode

Entity Model (Entity Mode)   
DAL (Data Access Layer)   
IDAL (Interface Layer)   
DALFactory (Class Factory)   
BLL (Business Logic Layer)   
BOF Business Object Framework   
SOA Service Orient Architecture Service Oriented Design   
EMF Eclipse Model Framework Eclipse Modeling frame

![](https://img2020.cnblogs.com/blog/13318/202109/13318-20210927093330053-1240377543.png)

# 一、简单Java对象

1️⃣PO `persistent object`

持久对象。与数据库里表字段一一对应。PO是一些属性，以及set和get方法组成。一般情况下，一个表对应一个PO，直接与操作数据库的crud相关。

2️⃣VO `view object`/`value object`

表现层对象。通常用于业务层之间的数据传递，和PO一样也是仅仅包含数据而已。但应是抽象出的业务对象，可以和表对应，也可以不。这根据业务的需要而定。对于页面上要展示的对象，可以封装一个VO对象，将所需数据封装进去。

3️⃣DTO `data trasfer object`

数据传输对象。主要用于远程调用等需要大量传输对象的地方。

比如一张表有 100 个字段，那么对应的 PO 就有 100 个属性。 但是界面上只要显示 10 个字段， 客户端用 WEB service 来获取数据，没有必要把整个 PO 对象传递到客户端，

这时就可以用只有这 10 个属性的 DTO 来传递结果到客户端，这样也不会暴露服务端表结构 . 到达客户端以后，如果用这个对象来对应界面显示，那此时它的身份就转为 VO。

4️⃣POJO `plain ordinary java object`/`pure old java object`

简单无规则 java 对象，纯的传统意义的 java 对象。

# 二、复杂Java对象

1️⃣BO/DO `bussiness object`/`Domain Object`

业务对象、域对象。封装业务逻辑的 Java 对象，通过调用 DAO 方法，结合 PO，VO 进行业务操作。一个BO对象可以包括多个PO对象。如常见的工作简历例子为例，简历可以理解为一个BO，简历又包括工作经历，学习经历等，这些可以理解为一个个的PO，由多个PO组成BO。

2️⃣DAO `data access object`

数据访问对象。此对象用于访问数据库。通常和 PO 结合使用，DAO 中包含了各种数据库的操作方法。通过它的方法，结合 PO 对数据库进行相关的操作。夹在业务逻辑与数据库资源中间。

  
  
  

REF

https://www.cnblogs.com/zlw-xf/p/9298393.html

https://www.jianshu.com/p/e8ee605662ea

https://blog.birost.com/a?ID=00400-6d6a3323-a504-457a-8376-2b6ecca15b10

https://blog.krybot.com/a?ID=00450-b981bea2-2a2b-4a79-b69e-80a9669b8e53

https://blog.karthisoftek.com/a?ID=00050-9991c590-f910-40e6-a4c0-6f81733d6b00

https://www.jianshu.com/p/d9cfd1a85068

https://stackoverflow.com/questions/1612334/difference-between-dto-vo-pojo-javabeans