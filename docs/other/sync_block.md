一切事物都是从同步阻塞开始，慢慢发展成异步非阻塞

只要不想阻塞，就有异步的需求，因此就要实现并发，在计算机就由多进程、多线程来具体实现。

## **同步、异步**

---

同步、异步是一对，关注的是消息通知机制，一次完成是同步，两次完成是异步

- 同步: 打电话（一次性把事情说完）

- 异步: 发短信说事，等待回复（收到回复才算是把事情说完，这里是两次）

同步与异步是对应的，它们是线程之间的关系，两个线程之间要么是同步的，要么是异步的。

## **阻塞、非阻塞**

---

阻塞、非阻塞是一对，关注的是程序在等待调用结果（消息，返回值）时的状态

- 阻塞: 只能做一件事

- 非阻塞: 做一件事同时可以做其他事

阻塞与非阻塞是对同一个线程来说的，在某个时刻，线程要么处于阻塞，要么处于非阻塞。

阻塞是使用同步机制的结果，非阻塞则是使用异步机制的结果。

## **同步异步和阻塞非阻塞的组合**

---

同步、异步，与阻塞、非阻塞，没有必然联系

组合例子：

- 同步阻塞: 打电话时候，其他任何事都做不了（比如不能同时看孩子）。效率低下

- 同步非阻塞: 打电话时候，时不时看看孩子有没有危险。不停切换，效率低下

- 异步阻塞: 发短信，在收到短信回复之前什么事都做不了（比如不能同时看孩子）。这种情形不常见

- 异步非阻塞: 发短信，然后陪孩子玩，收到短信回复通知后再去看手机。效率最高

编码应用，指的都是用户代码逻辑，而非底层的实现：

- 同步阻塞：最普遍最常见，符合思维，容易编写，容易调试

- 同步非阻塞：自己的逻辑是同步，但由于使用了某个模块，而这个模块是实现异步的，所以立刻返回，于是从自己的代码逻辑角度看没有阻塞。

	!!! warning ""
		如果底层没有实现异步，那么上层的用户逻辑就不可能实现非阻塞。比如nodejs都说是单线程、异步、非阻塞I/O，其实是因为nodejs底层实现了异步（比如多线程之类的），但没有开放异步功能（比如多线程）给用户使用。

## **单线程、多线程**

---

- 单线程: 单线程只能实现同步，无法实现异步

- 多线程: 是实现异步的最常用手段，实现异步还有多进程、协程等方法

## **进程、线程、协程**

---

- 进程:

	- 概念: 资源分配的基本单位

	- 通信: 进程之间的通信只能通过进程通信的方式进行

	- 多进程: 拷贝，使用fork()，生成子进程。每个进程拥有独立的地址空间（代码段、堆栈段、数据段）

- 线程:

	- 概念: 调度运行的最小单位

	- 通信: 同一进程中的线程共享数据（比如全局变量，静态变量）

	- 多线程: 同一个进程中的线程，它们之间共享大部分数据，使用相同的地址空间。当然线程是拥有自己的局部变量和堆栈（注意不是堆）

- 协程:

	- 概念: 用户态模拟进程线程的切换的具体实现，并非OS内核提供的功能。由程序员主动控制协程之间的切换。

	- 通信: 不要通过共享内存来通信，而应该通过通信来共享内存。

		!!! note ""
			golang提供一种基于消息机制而非共享内存的通信模型。消息机制认为每个并发单元都是自包含的独立个体，并且拥有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。

	- 应用场景: 经典的web应用场景就是推送的并发能力

!!! warning
	实际上协程是最早出现的，那个时候计算机的能力还很单一，不具备多进程、多线程功能，是由程序员手工控制多任务之间的调度。后来进程出现了，由于性能瓶颈，出现了线程。一直到这个时候协程都是不温不火的，但后来由于单cpu的计算能力已不成为瓶颈，而是io成为瓶颈，因此协程又逐渐火了起来。因为协程的切换开销小，适用于io密集型高并发场景。

效率（即切换开销谁最小）比较：协程 > 多线程 > 多进程

编程难度比较: 协程 > 多进程 > 多线程。因为多线程要注意同步和互斥（线程安全），否则容易造成死锁。

??? note "经典讲解（网上摘取）"
	- 进程是cpu资源分配的最小单位，线程是cpu调度的最小单位。以前进程既是资源分配也是调度的最小单位，后来为了更合理的使用cpu(实际上是cpu性能越来越好)，才将资源分配和调度分开，就有了线程。线程是建立在进程的基础上的一次程序运行单位。

	- 如果说多进程对于多CPU，多线程对应多核CPU，那么事件驱动和协程则是在充分挖掘不断提高性能的单核CPU的潜力。

	- 在操作系统中，进程拥有自己独立的内存空间，多个进程同时运行不会相互干扰，但是进程之间的通信比较麻烦；线程拥有独立的栈但共享内存，因此数据共享比较容易，但是多线程中需要利用加锁来进行访问控制：这是个非常头痛的问题，不加锁非常容易导致数据的错误，加锁容易出现死锁。多线程在多个核上同时运行，程序员根本无法控制程序具体的运行过程，难以调试。而且线程的切换经常需要深入到内核，因此线程的切换代价还是比较大的。协程coroutine拥有自己的栈和局部变量，相互之间共享全局变量。任何时候只有一个协程在真正运行，程序员能够控制协程的切换和运行，因此协程的程序相比多线程编程来说轻松很多。由于协程是用户级线程，因为协程的切换代价很小。

	- 程序员能够控制协程的切换，这句话需要认真理解下。程序员通过yield让协程在空闲（比如等待io，网络数据未到达）时放弃执行权，通过resume调度协程运行。协程一旦开始运行就不会结束，直到遇到yield交出执行权。Yield和resume这一对控制可以比较方便地实现程序之间的“等待”需求，即“异步逻辑”。总结起来，就是协程可以比较方便地实现用同步的方式编写异步的逻辑。

	- 从概念上讲，协程，又称微线程、纤程，英文名 Coroutine。它类似于子例程，但执行过程中，在子例程内部可中断，然后转而执行别的子例程，在适当的时候再返回来接着执行。以一个 Web 请求的处理为例，当用户发起请求时，我们需要从后端数据库中提取所需的数据，这个操作可能会很慢。这时，如果是协程，则可以在此先中断，转而让别人先执行，待数据库的查询返回之后再返回来继续完成数据的组装并最终返回给用户。这样就感觉和谐了许多。

	- 我们前面说过，协程的思想本质上就是控制流的主动让出和恢复机制。在现代语言里，可以实现协程思想的方法很多，这些实现间并无高下之分，所区别的就是是否适合应用场景。理解这一点，我们对于各种协程的分类，如半对称/对称协程，有栈与无栈协程等具体实现就能提纲挈领，无需在实现细节上纠结。

	- 总的来说，协程为协同任务提供了一种运行时抽象。这种抽象非常适合于协同多任务调度和数据流处理。在现代操作系统和编程语言中，因为用户态线程切换代价比内核态线程小，协程成为了一种轻量级的多任务模型。我们无法预测未来，但是可以看到，协程已经成为许多擅长数据处理的语言的一级对象。随着计算机并行性能的提升，用户态任务调度已经成为一种标准的多任务模型。在这样的大趋势下，协程这个简单且有效的模型就显得更加引人注目。

	- 以现代眼光来看，高级语言编译器实际上是多个步骤组合而成：词法解析，语法解析，语法树构建，以及优化和目标代码生成等等。编译实质上就是从源程序出发，依次将这些步骤的输出作为下一步的输入，最终输出目标代码。在现代计算机上实现这种管道式的架构毫无困难：只需要依次运行，中间结果存为中间文件或放入内存即可。GCC 和 Clang 编译器，以及 ANTLR 构建的编译器，都遵循这样的设计。而在 Conway 的设计里，词法和语法解析不再是两个独立运行的步骤，而是交织在一起。编译器的控制流在词法和语法解析之间来回切换：当词法模块读入足够多的 token 时，控制流交给语法分析；当语法分析消化完所有 token 后，控制流交给词法分析。词法和语法分别独立维护自身的运行状态。Conway 构建的这种协同工作机制，需要参与者“让出 (yield)”控制流时，记住自身状态，以便在控制流返回时能够从上次让出的位置恢复(resume)执行。简言之，协程的全部精神就在于控制流的主动让出和恢复。我们熟悉的子过程调用可以看作在返回时让出控制流的一种特殊的协程，其内部状态在返回时被丢弃了，因此不存在“恢复”这个操作。

## **并发、并行、多任务**

---

- 并发: (宏观的)，逻辑上的同时运行，但同一时间底层可能只有一个任务在运行

- 并行: (微观的)，真正的同时运行，通过多核实现，任务分布在不同核上，也可以是分布在不同服务器上。

- 多任务: 只要逻辑上能运行多个程序，就算是多任务，无论是否真正并行。

单核cpu就是并发，多核cpu才有可能并行，为什么说是有可能而不是绝对，因为系统上运行的程序数量必定比cpu数量多

## **cpu密集型和io密集型**

---

- io密集型: 包含网络io和磁盘io等，就是一个执行很久的操作不是由本机cpu进行操作而是发送一个请求然后由远程服务器操作，或者是发送一个读写磁盘的请求然后由磁盘硬件DMA进行操作。总之就是cpu发出一个请求，然后等待结果。如果一个业务的这种类型的操作比较多，就称这个业务为io密集型业务

- cpu密集型: 一个操作由本机cpu进行运算。如果一个业务的cpu运算量比较大，而io操作比较少，则称这个业务为cpu密集型业务

多核处理器可以实现cpu运算的并行，提高运算能力，这个时候用多进程或者多线程就适用于cpu密集型。而协程不适合，因为协程是用户态的技术，只能单核，因此并不能提高cpu并行运算能力，也因此只适合io密集型。多进程和多线程除了在多核时候适应cpu密集型，也可以适用io密集型，尤其是单核时候，只不过协程的切换开销比进程和线程都小

!!! note "经典讲解（网上摘取）"
	总的说来，当单核cpu性能提升，cpu不在成为性能瓶颈时，采用异步server能够简化编程模型，也能提高IO密集型应用的性能。

	在 I/O 密集型的应用中，CPU 可能总是苦苦等待着 I/O 操作的完成。如果是一个提供 Web 的服务的话，也就意味着一个线程会因为 I/O 阻塞而无法快速的对其他请求进行响应。势必也造成一种资源浪费和效率低下。在这种时候，协程的价值就体现了出来，例如基于 Python 语言实现的 tornado Web 框架就是基于此原理。

## **协同式调度、抢占式调度**

---

因为我们所熟悉的桌面操作系统，都是从协同式调度（如 Windows 3.2， Mac OS 9 等）过渡到抢占式多任务系统的。实际上，调度方式并无高下，完全取决于应用场景。抢占式系统允许操作系统剥夺进程执行权限，抢占控制流，因而天然适合服务器和图形操作系统，因为调度器可以优先保证对用户交互和网络事件的快速响应。当年 Windows 95 刚刚推出的时候，抢占式多任务就被作为一大买点大加宣传。协同式调度则等到进程时间片用完或系统调用时转移执行权限，因此适合实时或分时等等对运行时间有保障的系统。

## **事件驱动、观察者、生产者消费者、订阅发布、服务发现**

---

下面几个术语，其实原理都差不多，都是一种主动式的操作，只是有的名词早出现，有的名词晚出现。

- 事件驱动: 是一种宏观的经典模式，等待消息做事，就是事件驱动，下面的观察者、生产者消费者、订阅发布、服务发现都只是事件驱动的具体实现模型。（个人理解）

	??? note "经典讲解（网上摘取）"
		现下流行的异步server都是基于事件驱动的（如nginx）。事件驱动简化了编程模型，很好地解决了多线程难于编程，难于调试的问题。异步事件驱动模型中，把会导致阻塞的操作转化为一个异步操作，主线程负责发起这个异步操作，并处理这个异步操作的结果。由于所有阻塞的操作都转化为异步操作，理论上主线程的大部分时间都是在处理实际的计算任务，少了多线程的调度时间，所以这种模型的性能通常会比较好。

- 观察者模式: 我认为这是事件驱动的具体模式，而下面的生产者消费者、订阅发布、服务发现，又是观察者模式的具体实现方式。

	??? note "经典讲解（网上摘取）"
		有时又被称为发布（publish ）-订阅（Subscribe）模式、模型-视图（View）模式、源-收听者(Listener)模式或从属者模式。属于行为型模式的一种，它定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。这个主题对象在状态变化时，会通知所有的观察者对象，使他们能够自动更新自己。

- 生产消费者模式: 指的是由生产者将数据源源不断推送到消息中心，由不同的消费者从消息中心取出数据做自己的处理，在同一类别下，所有消费者拿到的都是同样的数据

- 订阅发布模式: 本质上也是一种生产消费者模式，不同的是，由订阅者首先向消息中心指定自己对哪些数据感兴趣，发布者推送的数据经过消息中心后，每个订阅者拿到的仅仅是自己感兴趣的一组数据。

- 服务注册与发现: 微服务架构由一个服务注册中心和围绕该注册中心的一堆服务组成。每个服务将自己的地址和其他元信息注册到注册中心（zookeeper或etcd等），有依赖它的其他服务会实时发现（watch机制）变化。

生产消费者模式和订阅发布模式是使用消息中间件时最常用的，用于功能解耦和分布式系统间的消息通信。