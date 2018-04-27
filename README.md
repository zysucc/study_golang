# study_golang

开始学习golang语言啦


go语言特性之一   type  test struct {} 结构体中 可以自定义工厂函数

    func  CreateNode (value int) *test{
      return &test{value: value}
    }  
<div>
    在一般的编程语言中是不支持局部变量在外部使用的,这取决你变量存放的内存空间的地址
# 栈内存
     栈内存--首先是一片内存区域，存储的都是局部变量，凡是定义在方法中的都是局部变量
<br/>（方法外的是全局变量），for循环内部定义的也是局部变量，是先加载函数才能进行局
<br/>部变量的定义，所以方法先进栈，然后再定义变量，变量有自己的作用域，一旦离开作
<br/>用域，变量就会被释放。栈内存的更新速度很快，因为局部变量的生命周期都很短
# 堆内存
     堆内存--存储的是数组和对象（其实数组就是对象），凡是new建立的都是在堆中，堆中
<br/>存放的都是实体（对象），实体用于封装数据，而且是封装多个（实体的多个属性），
<br/>如果一个数据消失，这个实体也没有消失，还可以用，所以堆是不会随时释放的，但是栈
<br/>不一样，栈里存放的都是单个变量，变量被释放了，那就没有了。堆里的实体虽然不会被
<br/>释放，但是会被当成垃圾，Java有垃圾回收机制不定时的收取。
</div>


