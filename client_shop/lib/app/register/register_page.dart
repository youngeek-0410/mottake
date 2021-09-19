import 'package:client_shop/app/providers.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class RegisterPage extends ConsumerStatefulWidget {
  const RegisterPage({Key? key}) : super(key: key);

  @override
  _RegisterPageState createState() => _RegisterPageState();
}

class _RegisterPageState extends ConsumerState<RegisterPage> {
  final _shopNameController = TextEditingController();
  final _shopDescriptionController = TextEditingController();
  final _shopAddressController = TextEditingController();
  final _shopSalesGoalController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _registerForm(),
    );
  }

  Widget _registerForm() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        // Username TextField
        TextField(
          controller: _shopNameController,
          decoration: const InputDecoration(
              icon: Icon(Icons.store), labelText: 'ShopName'),
        ),

        // Email TextField
        TextField(
          controller: _shopDescriptionController,
          decoration: const InputDecoration(
              icon: Icon(Icons.note), labelText: 'ShopDescription'),
        ),

        // Password TextField
        TextField(
          controller: _shopAddressController,
          decoration: const InputDecoration(
              icon: Icon(Icons.sort_by_alpha), labelText: 'ShopAddress'),
        ),

        TextField(
          controller: _shopSalesGoalController,
          decoration: const InputDecoration(
              icon: Icon(Icons.money), labelText: 'SalesGoal'),
          keyboardType: TextInputType.number,
          inputFormatters: [FilteringTextInputFormatter.digitsOnly],
        ),

        // Sign Up Button
        ElevatedButton(
          onPressed: () {
            _register();
          },
          child: const Text("Register"),
        )
      ],
    );
  }

  void _register() {
    final name = _shopNameController.text.trim();
    final description = _shopDescriptionController.text.trim();
    final address = _shopAddressController.text.trim();
    final salesGoal = int.parse(_shopSalesGoalController.text.trim());
    final auth = ref.read(authServiceProvider);
    auth.register(name, description, address, salesGoal);
  }
}
