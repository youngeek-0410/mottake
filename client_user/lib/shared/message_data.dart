import 'package:client_user/services/database.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final messageProvider = FutureProvider.autoDispose<SampleResponse>((ref) async {
  final message = await getMessage();
  return message;
});
