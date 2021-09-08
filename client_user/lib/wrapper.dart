import 'package:client_user/pages/home_page.dart';
import 'package:client_user/pages/login_page.dart';
import 'package:client_user/shared/auth_state_provider.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class App extends ConsumerWidget {
  @override
  Widget build(BuildContext context, ScopedReader watch) {
    return watch(authStateProvider).when(
        data: (user) {
          if (user == null) {
            return LoginPage();
          } else {
            return HomePage();
          }
        },
        loading: () => Scaffold(
              body: Center(
                child: CircularProgressIndicator(),
              ),
            ),
        error: (error, stackTrace) => Text('${error.toString()}'));
  }
}
