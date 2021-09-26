import 'dart:convert';

import 'package:client_user/app/config/exceptions.dart';
import 'package:client_user/app/home/models/customer.dart';
import 'package:client_user/app/home/models/shop.dart';
import 'package:client_user/app/home/models/menu.dart';
import 'package:client_user/constants/urls.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:client_user/app/home/map/map_page.dart';
import 'package:http/http.dart' as http;
import 'package:location/location.dart';

class Database {
  Database({required this.user});
  final User user;

  Future<Map<String, String>> getHeader() async {
    final header = {"Authorization": "Bearer " + await user.getIdToken()};
    return header;
  }

  Future<Customer> getCustomer() async {
    try {
      final url = URLs.baseURL + "/user";
      final header = await getHeader();
      final response = await http.get(Uri.parse(url), headers: header);
      if (response.statusCode != 200) {
        throw APIException(response.body.toString());
      }
      final customer = Customer.fromJson(jsonDecode(response.body));
      return customer;
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  Future<void> deleteCustomer() async {
    try {
      final url = URLs.baseURL + "/user";
      final header = await getHeader();
      final response = await http.delete(Uri.parse(url), headers: header);
      if (response.statusCode != 204) {
        throw APIException(response.body);
      }
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  Future<Shops?> searchShop(LocationData? location) async {
    try {
      final url = URLs.baseURL +
          "/shop/search?latitude=" +
          location!.latitude.toString() +
          "&longitude=" +
          location.longitude.toString();
      final header = await getHeader();
      final response = await http.get(Uri.parse(url), headers: header);
      if (response.statusCode != 200) {
        throw APIException(response.body.toString());
      }
      final shop = Shops.fromJson(jsonDecode(response.body));
      return shop;
    } catch (e) {
      //throw APIException(e.toString());
      return null;
    }
  }

  Future<MenuList> getMenus(String shopUID) async {
    try {
      // final url = URLs.baseURL + "/shop/" + user.uid + "/menu";
      final url = URLs.baseURL + "/shop/" + shopUID + "/menu";
      final header = await getHeader();
      final response = await http.get(Uri.parse(url), headers: header);
      if (response.statusCode != 200) {
        throw APIException(response.body);
      }
      final menuList = MenuList.fromJson(jsonDecode(response.body));
      return menuList;
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  String generateQRcode(Map<int, int> shoppingcourt) {
    final uid = user.uid;
    List<dynamic> purchases = [];
    shoppingcourt.forEach((key, value) {
      final purchase = {"menu_id": key, "number": value};
      purchases.add(purchase);
    });
    Map<String, dynamic> result = {"customer_uid": uid, "purchases": purchases};
    return jsonEncode(result);
  }
}
