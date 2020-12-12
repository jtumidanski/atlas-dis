package com.atlas.dis.database.transformer;

import com.atlas.dis.entity.MonsterDrop;
import com.atlas.dis.model.MonsterDropData;
import transformer.SqlTransformer;

public class MonsterDropTransformer implements SqlTransformer<MonsterDropData, MonsterDrop> {
   @Override
   public MonsterDropData transform(MonsterDrop monsterDrop) {
      return new MonsterDropData(monsterDrop.getId(), monsterDrop.getMonsterId(), monsterDrop.getItemId(),
            monsterDrop.getMaximumQuantity(), monsterDrop.getMinimumQuantity(), monsterDrop.getChance());
   }
}
