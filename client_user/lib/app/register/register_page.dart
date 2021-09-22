import 'package:client_user/app/providers.dart';
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
  final _customerNameController = TextEditingController();
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
          controller: _customerNameController,
          decoration: const InputDecoration(
              icon: Icon(Icons.person), labelText: 'Name'),
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
    final name = _customerNameController.text.trim();
    final auth = ref.read(authServiceProvider);
    auth.register(name);
  }
}
