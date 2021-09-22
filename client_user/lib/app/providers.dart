import 'package:client_user/service/auth_service.dart';
import 'package:client_user/service/database.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final firebaseAuthProvider =
    Provider<FirebaseAuth>((ref) => FirebaseAuth.instance);

final authServiceProvider = Provider((ref) {
  final firebaseAuth = ref.watch(firebaseAuthProvider);
  return AuthService(firebaseAuth: firebaseAuth);
});

final authStateProvider = StreamProvider(
    (ref) => ref.watch(authServiceProvider).authStateController.stream);

final databaseProvider = Provider((ref) {
  final authService = ref.watch(authServiceProvider);
  if (authService.user != null) {
    return Database(user: authService.user!);
  } else {
    return null;
  }
});
