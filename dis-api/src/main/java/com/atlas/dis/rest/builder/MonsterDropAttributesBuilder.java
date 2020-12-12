package com.atlas.dis.rest.builder;

import builder.AttributeResultBuilder;
import builder.RecordBuilder;
import com.atlas.dis.rest.attribute.MonsterDropAttributes;

public class MonsterDropAttributesBuilder extends RecordBuilder<MonsterDropAttributes, MonsterDropAttributesBuilder> implements AttributeResultBuilder {
   private static final String MONSTER_ID = "MONSTER_ID";

   private static final String ITEM_ID = "ITEM_ID";

   private static final String MAXIMUM_QUANTITY = "MAXIMUM_QUANTITY";

   private static final String MINIMUM_QUANTITY = "MINIMUM_QUANTITY";

   private static final String CHANCE = "CHANCE";


   @Override
   public MonsterDropAttributes construct() {
      return new MonsterDropAttributes(get(MONSTER_ID), get(ITEM_ID), get(MAXIMUM_QUANTITY), get(MINIMUM_QUANTITY), get(CHANCE));
   }

   @Override
   public MonsterDropAttributesBuilder getThis() {
      return this;
   }

   public MonsterDropAttributesBuilder setMonsterId(Integer monsterId) {
      return set(MONSTER_ID, monsterId);
   }

   public MonsterDropAttributesBuilder setItemId(Integer itemId) {
      return set(ITEM_ID, itemId);
   }

   public MonsterDropAttributesBuilder setMaximumQuantity(Integer maximumQuantity) {
      return set(MAXIMUM_QUANTITY, maximumQuantity);
   }

   public MonsterDropAttributesBuilder setMinimumQuantity(Integer minimumQuantity) {
      return set(MINIMUM_QUANTITY, minimumQuantity);
   }

   public MonsterDropAttributesBuilder setChance(Integer chance) {
      return set(CHANCE, chance);
   }

}
