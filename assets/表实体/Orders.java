import java.util.*;

/**
 * 定单实体
 */


public class Orders{
	private Integer id;				//定单编号
	private Integer uid;			//用户编号
	private String pay;				//付款方式
	private String carry;			//邮寄方式
	private String address;			//邮寄地址
	private Date orderDate;			//定单生效日期
	private String bz;				//备注信息
	private boolean enforce;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setUid(Integer uid) {
		this.uid = uid; 
	}

	public void setPay(String pay) {
		this.pay = pay; 
	}

	public void setCarry(String carry) {
		this.carry = carry; 
	}

	public void setAddress(String address) {
		this.address = address; 
	}

	public void setOrderDate(Date orderDate) {
		this.orderDate = orderDate; 
	}

	public void setBz(String bz) {
		this.bz = bz; 
	}

	public void setEnforce(boolean enforce) {
		this.enforce = enforce; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public Integer getUid() {
		return (this.uid); 
	}

	public String getPay() {
		return (this.pay); 
	}

	public String getCarry() {
		return (this.carry); 
	}

	public String getAddress() {
		return (this.address); 
	}

	public Date getOrderDate() {
		return (this.orderDate); 
	}

	public String getBz() {
		return (this.bz); 
	}

	public boolean getEnforce() {
		return (this.enforce); 
	}		//是否已执行
	
	
	
	
}