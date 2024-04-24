/**
 * 投票实体
 */


public class Poll{
	private Integer id;				//投票项编号
	private String optionName;		//投票项
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
	}			//得票数
	
	
}