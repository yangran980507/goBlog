/**
 * 折扣实体(等级)
 */


public class Rebate{
	private Integer id;			//折扣编号
	private Float amount;		//折扣金额(等级金额)
	private Float rebate;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setAmount(Float amount) {
		this.amount = amount; 
	}

	public void setRebate(Float rebate) {
		this.rebate = rebate; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public Float getAmount() {
		return (this.amount); 
	}

	public Float getRebate() {
		return (this.rebate); 
	}		//享受折扣比率
	
	
}