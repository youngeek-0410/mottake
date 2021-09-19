import 'package:client_shop/app/providers.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_signin_button/button_list.dart';
import 'package:flutter_signin_button/button_view.dart';

class LoginPage extends ConsumerWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authService = ref.watch(authServiceProvider);
    return Scaffold(
      body: Center(
          child: Column(
        children: [
          SignInButton(Buttons.Google, onPressed: () {
            authService.signInWithGoogle();
          })
        ],
        mainAxisAlignment: MainAxisAlignment.center,
      )),
    );
  }
}
