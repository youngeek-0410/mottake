import 'package:client_shop/app/home/models/shop.dart';
import 'package:client_shop/app/home/shop/shop_edit_page.dart';
import 'package:client_shop/app/providers.dart';
import 'package:client_shop/constants/strings.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final shopProvider = FutureProvider<Shop>(
  (ref) => ref.read(databaseProvider)!.getShop(),
);

class ShopPage extends ConsumerWidget {
  const ShopPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final auth = ref.watch(authServiceProvider);
    final _shopProvider = ref.watch(shopProvider);

    return _shopProvider.when(data: (data) {
      return Scaffold(
        body: Center(
            child: Column(
          children: [
            ListView(
              shrinkWrap: true,
              children: [
                _listItem(data.name!, const Icon(Icons.store)),
                _listItem(data.address!, const Icon(Icons.sort_by_alpha)),
                _listItem(data.description!, const Icon(Icons.note)),
                _salesIndicator(data.sales!, data.salesGoal!)
              ],
            ),
            ElevatedButton(
              child: const Text(Strings.signOut),
              onPressed: auth.signOut,
            ),
            ElevatedButton(
                child: const Text(Strings.edit),
                onPressed: () {
                  Navigator.of(context).push(MaterialPageRoute(builder: (_) {
                    return ShopEditPage(data);
                  })).then((_) => ref.refresh(shopProvider));
                }),
          ],
        )),
      );
    }, loading: () {
      return const CircularProgressIndicator();
    }, error: (error, stackTrace) {
      return const CircularProgressIndicator();
    });
  }

  Widget _listItem(String title, Icon icon) {
    return ListTile(title: Text(title), leading: icon);
  }

  Widget _salesIndicator(int sales, int salesGoal) {
    return ListTile(
      leading: const Icon(Icons.money),
      title: Text("$sales / $salesGoal yen"),
      subtitle: LinearProgressIndicator(
        value: sales / salesGoal,
      ),
    );
  }
}
