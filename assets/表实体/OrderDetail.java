/**
 * 定单明细实体
 */


public class OrderDetail{
	private Integer id;				//定单明细编号
	private Integer orderId;		//对应定单编号
	private Integer bookId;			//图书编号
	private Integer quantity;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setOrderId(Integer orderId) {
		this.orderId = orderId; 
	}

	public void setBookId(Integer bookId) {
		this.bookId = bookId; 
	}

	public void setQuantity(Integer quantity) {
		this.quantity = quantity; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public Integer getOrderId() {
		return (this.orderId); 
	}

	public Integer getBookId() {
		return (this.bookId); 
	}

	public Integer getQuantity() {
		return (this.quantity); 
	}		//购买数量
	
	
}