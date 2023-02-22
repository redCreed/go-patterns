package main

// House 要构建的具体产品属性
type House struct {
	window       string //窗
	door         string //门
	swimmingPool string //游泳池
	floor        int    //楼层
}

// IBuilder 抽象建造者：它为创建一个产品 Product 对象的各个部件指定抽象接口，这个接口一般包括两类方法
//buildPartX() ：用于创建复杂对象的各个部件
//getResult() ：用于返回复杂对象
type IBuilder interface {
	setWindow()
	setDoor()
	setSwimmingPool()
	setFloor()
	getHouse() House
}

//别墅 新建具体建造者
type villaBuilder struct {
	//window       string //窗
	//door         string //门
	//swimmingPool string //游泳池
	//floor        int    //楼层
	House
}

func newVillaBuilder() *villaBuilder {
	return &villaBuilder{}
}

func (b *villaBuilder) setWindow() {
	b.window = "Wooden Window"
}

func (b *villaBuilder) setDoor() {
	b.door = "Wooden Door"
}

func (b *villaBuilder) setFloor() {
	b.floor = 3
}

func (b *villaBuilder) setSwimmingPool() {
	b.swimmingPool = "None"
}

func (b *villaBuilder) getHouse() House {
	return House{
		window:       b.window,
		door:         b.door,
		swimmingPool: b.swimmingPool,
		floor:        b.floor,
	}
}

// Director 负责安排负责对象的建造次序，比如先确定门、窗、楼层，再考虑是否需要加装泳池
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoor()
	d.builder.setWindow()
	d.builder.setFloor()
	d.builder.setSwimmingPool()
	return d.builder.getHouse()
}
