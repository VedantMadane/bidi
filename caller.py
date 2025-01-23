from opperai import Opper
from opperai.types import CallConfiguration

opper = Opper()

output, _ = opper.call(
        name="call_with_max_tokens",
        instructions="answer the following question",
        input="what are some uses of 42",
        configuration=CallConfiguration(
            model_parameters={
                "max_tokens": 10,
            }
        ),
    )
print(output)