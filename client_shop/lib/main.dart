import 'package:client_shop/app/home/home_page.dart';
import 'package:client_shop/app/login/login_page.dart';
import 'package:client_shop/app/providers.dart';
import 'package:client_shop/app/register/register_page.dart';
import 'package:client_shop/service/auth_service.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  runApp(const ProviderScope(child: App()));
}

class App extends StatelessWidget {
  const App({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: Locator(),
    );
  }
}

class Locator extends ConsumerWidget {
  const Locator({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, ScopedReader watch) {
    return watch(authStateProvider).when(data: (data) {
      if (data.authFlowStatus == AuthFlowStatus.googleSignIn) {
        return LoginPage();
      } else if (data.authFlowStatus == AuthFlowStatus.register) {
        return RegisterPage();
      } else {
        return HomePage();
      }
    }, loading: () {
      return const CircularProgressIndicator();
    }, error: (error, stackTrace) {
      return LoginPage();
    });
  }
}
