import 'package:client_user/services/auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class LoginPage extends ConsumerWidget {
  @override
  Widget build(BuildContext context, ScopedReader watch) {
    final authService = watch(authServiceProvider);
    return Scaffold(
        body: Center(
            child: Column(
      children: [
        ElevatedButton(
          onPressed: authService.signIn,
          child: Text('Sign in with google'),
        ),
      ],
      mainAxisAlignment: MainAxisAlignment.center,
    )));
  }
}
