/**
 * 图书评论实体
 */

public class Remark{
	private Integer id;			//图书评论编号
	private Integer bookId;		//图书编号
	private Integer uid;		//用户编号
	private String content;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setBookId(Integer bookId) {
		this.bookId = bookId; 
	}

	public void setUid(Integer uid) {
		this.uid = uid; 
	}

	public void setContent(String content) {
		this.content = content; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public Integer getBookId() {
		return (this.bookId); 
	}

	public Integer getUid() {
		return (this.uid); 
	}

	public String getContent() {
		return (this.content); 
	}		//评论内容
	
	
	
}