class Shop {
  String? name;
  String? address;
  String? description;
  int? salesGoal;
  Shop({this.name, this.address, this.description, this.salesGoal});

  factory Shop.fromJson(Map<String, dynamic> json) => Shop(
      name: json["name"],
      address: json["address"],
      description: json["description"],
      salesGoal: json["sales_goal"]);

  Map<String, dynamic> toJson() => {
        "name": name,
        "address": address,
        "description": description,
        "sales_goal": salesGoal
      };
}
