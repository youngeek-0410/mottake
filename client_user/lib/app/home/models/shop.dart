class Shop {
  String? uid;
  String? name;
  String? address;
  String? description;
  int? sales;
  int? salesGoal;
  double? latitude;
  double? longitude;
  Shop(
      {this.uid,
      this.name,
      this.address,
      this.description,
      this.sales,
      this.salesGoal,
      this.latitude,
      this.longitude});

  factory Shop.fromJson(Map<String, dynamic> json) => Shop(
      uid: json["uid"],
      name: json["name"],
      address: json["address"],
      description: json["description"],
      sales: json["sales"],
      salesGoal: json["sales_goal"],
      latitude: json["latitude"],
      longitude: json["longitude"]);

  Map<String, dynamic> toJson() => {
        "name": name,
        "address": address,
        "description": description,
        "sales": sales,
        "sales_goal": salesGoal,
        "latitude": latitude,
        "longitude": longitude
      };
}

class Shops {
  List<Shop>? shops;

  Shops(this.shops);

  factory Shops.fromJson(List<dynamic> parsedJson) {
    List<Shop> shops;
    shops = parsedJson.map((e) => Shop.fromJson(e)).toList();
    return Shops(shops);
  }
}
