/**
 * ͼ������ʵ��
 */

public class Remark{
	private Integer id;			//ͼ�����۱��
	private Integer bookId;		//ͼ����
	private Integer uid;		//�û����
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
	}		//��������
	
	
	
}