import 'dart:convert';

import 'package:client_shop/app/config/exceptions.dart';
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
        throw FetchException(response.body.toString());
      }
      final shop = Shop.fromJson(jsonDecode(response.body));
      return shop;
    } catch (e) {
      throw FetchException(e.toString());
    }
  }
}
