class FetchException implements Exception {
  final String message;
  FetchException(this.message);
  @override
  String toString() {
    return message;
  }
}
