/**
 * ͼ��ʵ��
 */



import java.util.*;


public class Books{
	private Integer id;				//ͼ����
	private String isbn;			//ͼ��ISBN
	private String bookName;		//����
	private Integer type;			//ͼ�����ͣ����ͱ����
	private String publisher;		//������
	private String author;			//����
	private String introduce;		//ͼ�����
	private Float price;			//ͼ��۸�
	private Date pdate;				//��������
	private String conver;			//��Ƥ���洢ͼƬ·��
	private Date intime;			//���ʱ��
	private boolean newBook;		//�Ƿ�Ϊ����
	private boolean commend;		//�Ƿ��Ƽ�
	private Integer quantity;		//�����
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
	}			//���۳���
	
	
}