/**
 * ������Ϣʵ��
 */



import java.util.*;

public class Bbs{
	private Integer id;			//������
	private String title;		//�������
	private String content;		//��������
	private Date showtime;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setTitle(String title) {
		this.title = title; 
	}

	public void setContent(String content) {
		this.content = content; 
	}

	public void setShowtime(Date showtime) {
		this.showtime = showtime; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public String getTitle() {
		return (this.title); 
	}

	public String getContent() {
		return (this.content); 
	}

	public Date getShowtime() {
		return (this.showtime); 
	}		//����ʱ��
	
	
	
}