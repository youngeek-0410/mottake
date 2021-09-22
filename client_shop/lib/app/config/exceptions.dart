class CustomException implements Exception {
  final String message;
  CustomException(this.message);
  @override
  String toString() {
    return message;
  }
}

class APIException extends CustomException {
  APIException(String message) : super(message);
}
