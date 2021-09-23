import 'dart:convert';

import 'package:client_user/app/config/exceptions.dart';
import 'package:client_user/app/home/models/customer.dart';
import 'package:client_user/app/home/models/shop.dart';
import 'package:client_user/constants/urls.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:http/http.dart' as http;

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

  Future<Shop> searchShop() async {
    try {
      final url = URLs.baseURL + "/shop/search";
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
}
