/**
 * ������ϸʵ��
 */


public class OrderDetail{
	private Integer id;				//������ϸ���
	private Integer orderId;		//��Ӧ�������
	private Integer bookId;			//ͼ����
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
	}		//��������
	
	
}