/**
 * 公告信息实体
 */



import java.util.*;

public class Bbs{
	private Integer id;			//公告编号
	private String title;		//公告标题
	private String content;		//公告正文
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
	}		//发布时间
	
	
	
}