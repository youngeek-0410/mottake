class Customer {
  String? name;
  Customer({this.name});

  factory Customer.fromJson(Map<String, dynamic> json) =>
      Customer(name: json["name"]);

  Map<String, dynamic> toJson() => {"name": name};
}
