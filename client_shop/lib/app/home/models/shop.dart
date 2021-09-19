class Shop {
  String? name;
  String? address;
  String? description;
  int? salesGoal;
  double? latitude;
  double? longitude;
  Shop(
      {this.name,
      this.address,
      this.description,
      this.salesGoal,
      this.latitude,
      this.longitude});

  factory Shop.fromJson(Map<String, dynamic> json) => Shop(
      name: json["name"],
      address: json["address"],
      description: json["description"],
      salesGoal: json["sales_goal"],
      latitude: json["latitude"],
      longitude: json["longitude"]);

  Map<String, dynamic> toJson() => {
        "name": name,
        "address": address,
        "description": description,
        "sales_goal": salesGoal,
        "latitude": latitude,
        "longitude": longitude
      };
}
