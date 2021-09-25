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
  final _formKey = GlobalKey<FormState>();
  String _name = "";
  String _description = "";
  String _address = "";
  int _salesGoal = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _registerForm(),
    );
  }

  Widget _registerForm() {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [_buildForm()],
      ),
    );
  }

  void _register() {
    if (_validateAndSaveForm()) {
      ref
          .read(authServiceProvider)
          .register(_name, _description, _address, _salesGoal);
    }
  }

  bool _validateAndSaveForm() {
    final form = _formKey.currentState!;
    if (form.validate()) {
      form.save();
      return true;
    }
    return false;
  }

  Widget _buildForm() {
    return Form(
      key: _formKey,
      child: Column(crossAxisAlignment: CrossAxisAlignment.stretch, children: [
        ..._buildFormChildren(),
        ElevatedButton(
            onPressed: () {
              _register();
            },
            child: const Text("Register")),
      ]),
    );
  }

  List<Widget> _buildFormChildren() {
    return [
      TextFormField(
        decoration: const InputDecoration(labelText: 'Name'),
        keyboardAppearance: Brightness.light,
        initialValue: _name,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _name = value!,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Description'),
        keyboardAppearance: Brightness.light,
        keyboardType: TextInputType.multiline,
        minLines: 4,
        maxLines: null,
        initialValue: _description,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _description = value!,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Address'),
        keyboardAppearance: Brightness.light,
        initialValue: _address,
        validator: (value) =>
            (value ?? '').isNotEmpty ? null : 'Name can\'t be empty',
        onSaved: (value) => _address = value!,
      ),
      TextFormField(
        decoration: const InputDecoration(labelText: 'Price'),
        keyboardAppearance: Brightness.light,
        initialValue: _salesGoal.toString(),
        keyboardType: const TextInputType.numberWithOptions(
          signed: false,
          decimal: false,
        ),
        validator: (value) {
          final parsed = int.tryParse(value ?? '') ?? 0;
          return parsed > 0 ? null : 'Price must be greater than 0 yen';
        },
        onSaved: (value) => _salesGoal = int.tryParse(value ?? '') ?? 0,
      ),
    ];
  }
}
