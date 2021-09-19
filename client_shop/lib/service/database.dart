import 'dart:convert';

import 'package:client_shop/app/config/exceptions.dart';
import 'package:client_shop/app/home/models/manu.dart';
import 'package:client_shop/app/home/models/shop.dart';
import 'package:client_shop/constants/urls.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:http/http.dart' as http;

class Database {
  Database({required this.user});
  final User user;

  Future<Map<String, String>> getHeader() async {
    final header = {"Authorization": "Bearer " + await user.getIdToken()};
    return header;
  }

  Future<Shop> getShop() async {
    try {
      final url = URLs.baseURL + "/shop/" + user.uid;
      final header = await getHeader();
      final response = await http.get(Uri.parse(url), headers: header);
      if (response.statusCode != 200) {
        throw APIException(response.body.toString());
      }
      final shop = Shop.fromJson(jsonDecode(response.body));
      return shop;
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  Future<MenuList> getMenus() async {
    try {
      final url = URLs.baseURL + "/shop/" + user.uid + "/menu";
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

  Future<void> patchMenu(Menu menu) async {
    try {
      final url = URLs.baseURL + "/management/shop/menu/" + menu.id.toString();
      final header = await getHeader();
      final response = await http.patch(Uri.parse(url),
          body: jsonEncode(menu.toJson()), headers: header);
      if (response.statusCode != 204) {
        throw APIException(response.body);
      }
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  Future<void> createMenu(Menu menu) async {
    try {
      final url = URLs.baseURL + "/management/shop/menu";
      final header = await getHeader();
      final response = await http.post(Uri.parse(url),
          body: jsonEncode(menu.toJson()), headers: header);
      if (response.statusCode != 201) {
        throw APIException(response.body);
      }
    } catch (e) {
      throw APIException(e.toString());
    }
  }

  Future<void> deleteMenu(int id) async {
    try {
      final url = URLs.baseURL + "/management/shop/menu/" + id.toString();
      final header = await getHeader();
      final response = await http.delete(Uri.parse(url), headers: header);
      if (response.statusCode != 204) {
        throw APIException(response.body);
      }
    } catch (e) {
      throw APIException(e.toString());
    }
  }
}
