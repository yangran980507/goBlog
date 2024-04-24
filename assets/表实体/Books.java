/**
 * 图书实体
 */



import java.util.*;


public class Books{
	private Integer id;				//图书编号
	private String isbn;			//图书ISBN
	private String bookName;		//书名
	private Integer type;			//图书类型，类型表外键
	private String publisher;		//出版社
	private String author;			//作者
	private String introduce;		//图书介绍
	private Float price;			//图书价格
	private Date pdate;				//出版日期
	private String conver;			//封皮，存储图片路径
	private Date intime;			//入库时间
	private boolean newBook;		//是否为新书
	private boolean commend;		//是否推荐
	private Integer quantity;		//库存量
	private Integer selled;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setIsbn(String isbn) {
		this.isbn = isbn; 
	}

	public void setBookName(String bookName) {
		this.bookName = bookName; 
	}

	public void setType(Integer type) {
		this.type = type; 
	}

	public void setPublisher(String publisher) {
		this.publisher = publisher; 
	}

	public void setAuthor(String author) {
		this.author = author; 
	}

	public void setIntroduce(String introduce) {
		this.introduce = introduce; 
	}

	public void setPrice(Float price) {
		this.price = price; 
	}

	public void setPdate(Date pdate) {
		this.pdate = pdate; 
	}

	public void setConver(String conver) {
		this.conver = conver; 
	}

	public void setIntime(Date intime) {
		this.intime = intime; 
	}

	public void setNewBook(boolean newBook) {
		this.newBook = newBook; 
	}

	public void setCommend(boolean commend) {
		this.commend = commend; 
	}

	public void setQuantity(Integer quantity) {
		this.quantity = quantity; 
	}

	public void setSelled(Integer selled) {
		this.selled = selled; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public String getIsbn() {
		return (this.isbn); 
	}

	public String getBookName() {
		return (this.bookName); 
	}

	public Integer getType() {
		return (this.type); 
	}

	public String getPublisher() {
		return (this.publisher); 
	}

	public String getAuthor() {
		return (this.author); 
	}

	public String getIntroduce() {
		return (this.introduce); 
	}

	public Float getPrice() {
		return (this.price); 
	}

	public Date getPdate() {
		return (this.pdate); 
	}

	public String getConver() {
		return (this.conver); 
	}

	public Date getIntime() {
		return (this.intime); 
	}

	public boolean getNewBook() {
		return (this.newBook); 
	}

	public boolean getCommend() {
		return (this.commend); 
	}

	public Integer getQuantity() {
		return (this.quantity); 
	}

	public Integer getSelled() {
		return (this.selled); 
	}			//已售出量
	
	
}