/**
 * ͶƱʵ��
 */


public class Poll{
	private Integer id;				//ͶƱ����
	private String optionName;		//ͶƱ��
	private Integer poll;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setOptionName(String optionName) {
		this.optionName = optionName; 
	}

	public void setPoll(Integer poll) {
		this.poll = poll; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public String getOptionName() {
		return (this.optionName); 
	}

	public Integer getPoll() {
		return (this.poll); 
	}			//��Ʊ��
	
	
}