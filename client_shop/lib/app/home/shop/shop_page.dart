import 'package:client_shop/app/home/models/shop.dart';
import 'package:client_shop/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final shopProvider = FutureProvider<Shop>(
  (ref) => ref.read(databaseProvider)!.getShop(),
);

class ShopPage extends ConsumerWidget {
  const ShopPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, ScopedReader watch) {
    final auth = watch(authServiceProvider);
    final _shopProvider = watch(shopProvider);
    return Scaffold(
        body: Center(
            child: Column(
      children: [
        _shopProvider.when(data: (data) {
          return ListView(
            shrinkWrap: true,
            children: [
              _listItem(data.name!, Icon(Icons.store)),
              _listItem(data.address!, Icon(Icons.sort_by_alpha)),
              _listItem(data.description!, Icon(Icons.note)),
              _listItem(data.salesGoal!.toString(), Icon(Icons.money))
            ],
          );
        }, loading: () {
          return const CircularProgressIndicator();
        }, error: (error, stackTrace) {
          return const CircularProgressIndicator();
        }),
        ElevatedButton(
          child: const Text("Sign out"),
          onPressed: auth.signOut,
        ),
      ],
    )));
  }

  Widget _listItem(String title, Icon icon) {
    return ListTile(title: Text(title), leading: icon);
  }
}
