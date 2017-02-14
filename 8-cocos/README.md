# Cocos

![Cocos2d-x 引擎系统结构](01-Cocos2d-x引擎系统结构.png)

### Cocos2d-x 内存
--

#### C++ 11 智能指针
- C++ 11 使用3种不同的智能指针: unique_ptr、shared_ptr、weak_ptr(属于模板类型)
```c++
int main(){
    unique_ptr<int> up1(new int(11));
    unique_ptr<int> up11 = up1; // 编译报错

    shared_ptr<int> up2(new int(22));
    weak_ptr<int> up3 = up2;
}
```
- 每个智能指针都重载了 "*" 运算符，可以使用 "*up1" 访问所分配的堆内存。
- 智能指针在析构或者调用reset成员时，都可能释放其所拥有的堆内存

>区别:
>    unique_ptr指针不能与其他智能指针共享所指对象的内存，如"up1"赋值给 "up11" 将导致编译报错。但是，可以通过标准库的 move 函数来转移 unique_ptr 指针对象的 "所有权"，一旦转移成功，原来的 unique_ptr 指针就失去了对对象的所有权，再使用该指针会导致运行时错误。
>    多个shared_ptr指针可以共享同一堆分配对象的内存，它在实现采用引用计数。即使一个shared_ptr指针放弃了所有权(调用reset成员或离开其作用域)，也不会影响其他智能指针对象。只有所有引用计数归0，才会真正释放所占有的堆内存。
>    weak_ptr指针可以用来指向shared_ptr指针分配的对象内存，但不拥有该内存。可以使用其lock成员来访问其指向内存的一个shared_ptr对象，
当其所指向的内存无效时，返回指针空值(nullptr)。weak_ptr指针通常可以用来验证shared_ptr指针的有效性。


#### Cocos 为什么不用智能指针 shared_ptr
- 智能指针有比较大的性能损失。
    shared_ptr为了保证线程安全，必须使用一定形式的互斥锁来保证所有线程访问时其引用计数正确。
    这种性能损失对一般的应用程序没有问题，对游戏这种实时性要求高的应用程序却是不可接受的，游戏需要更简单的内存管理模型。
- 虽然智能指针能进行有效的堆内存管理，但仍需要显式的声明智能指针。
    ```c++
    // 例如，创建一个Node的代码:
    shared_ptr<Node *> node(new Node());
    // 在需要引用的地方，一般应该使用weak_ptr指针，否则在Node被移除时，还要手动减少shared_ptr指针的引用计数
    weak_ptr<Node*> refNode = node;
    ```

#### 垃圾回收
- 基于引用计数：使用系统记录的一个对象被引用的次数，当对象被引用的次数位0时，该对象即被视作垃圾而被回收。
- 基于跟踪处理:先产生跟踪对象的关系图，再进行垃圾回收。
    其算法是首先将程序中正在使用的对象视为根对象，从根对象开始查找它们所引用的堆空间，并在这些堆空间上做标记。
    做完标记后，所有未被标记的对象即被视作垃圾，会在第二阶段被清理。在第二阶段可以使用不同的方式进行清理，直接清理可能会产生大量的内存碎片。
    清理方法是对正在使用的对象进行移动或者复制，从而减少内存碎片的产生。

#### Cocos2d-x 内存管理机制
- 引用计数
Cocos2d-x中的所有对象几乎都继承自Ref基类。Ref基类主要的职责就是对对象进行引用计数管理
```c++
    class CC_DLL Ref
    {
    public:
        void retatin();
        void release();
        Ref* autorelease();
        unsigned int getReferenceCount() const;
    protected:
        Ref();
        // count of references
        unsigned int _referenceCount;
        friend class AutoreleasePool;

    }
```






























END