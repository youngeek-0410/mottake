import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ShopPage extends ConsumerWidget {
  const ShopPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, ScopedReader watch) {
    final auth = watch(authServiceProvider);
    return Scaffold(
        body: Center(
      child: ElevatedButton(
        child: const Text("Sign out"),
        onPressed: auth.signOut,
      ),
    ));
  }
}
