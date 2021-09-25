class Menu {
  int? id;
  String? shopUid;
  String? name;
  int? price;

  Menu({this.id, this.shopUid, this.name, this.price});

  factory Menu.fromJson(Map<String, dynamic> json) => Menu(
      id: json["id"],
      shopUid: json["shop_uid"],
      name: json["name"],
      price: json["price"]);

  Map<String, dynamic> toJson() =>
      {"id": id, "shop_uid": shopUid, "name": name, "price": price};
}

class MenuList {
  List<Menu>? menus;

  MenuList(this.menus);

  factory MenuList.fromJson(List<dynamic> parsedJson) {
    List<Menu> menus;
    menus = parsedJson.map((e) => Menu.fromJson(e)).toList();
    return MenuList(menus);
  }
}
