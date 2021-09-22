import 'dart:convert';

import 'package:client_user/app/config/exceptions.dart';
import 'package:client_user/app/home/models/customer.dart';
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
        throw FetchException(response.body.toString());
      }
      final customer = Customer.fromJson(jsonDecode(response.body));
      return customer;
    } catch (e) {
      throw FetchException(e.toString());
    }
  }
}
