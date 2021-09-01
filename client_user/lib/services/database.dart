import 'dart:convert';
import 'package:http/http.dart' as http;

class SampleResponse {
  final String message;
  SampleResponse(this.message);
  SampleResponse.fromJson(Map<String, dynamic> json)
      : message = json['message'];
  Map<String, dynamic> toJson() => {
        'message': message,
      };
}

Future<SampleResponse> getMessage() async {
  final url = 'http://10.0.2.2';
  final response = await http.get(Uri.parse(url));
  final jsonBody = json.decode(response.body);
  return SampleResponse.fromJson(jsonBody);
}
