class Receipt {
  String? customerUid;
  List<Purchase>? purchases;

  Receipt({this.customerUid, this.purchases});

  factory Receipt.fromJson(Map<String, dynamic> json) {
    final customerUid = json["customer_uid"];
    final purchases = List<dynamic>.from(json["purchases"])
        .map((e) => Purchase.fromJson(e))
        .toList();
    return Receipt(customerUid: customerUid, purchases: purchases);
  }

  Map<String, dynamic> toJson() =>
      {"customer_uid": customerUid, "purchases": purchases};
}

class Purchase {
  int? menuId;
  int? number;

  Purchase({this.menuId, this.number});
  factory Purchase.fromJson(Map<String, dynamic> json) =>
      Purchase(menuId: json["menu_id"], number: json["number"]);

  Map<String, dynamic> toJson() => {"menu_id": menuId, "number": number};
}
